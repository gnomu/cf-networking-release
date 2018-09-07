package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"policy-server/handlers"
	"policy-server/handlers/fakes"
	"policy-server/uaa_client"

	"policy-server/store"
	storeFakes "policy-server/store/fakes"

	"bytes"
	"errors"

	"code.cloudfoundry.org/cf-networking-helpers/httperror"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type failingReader struct {
}

func (f *failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("can't do it")
}

func (f *failingReader) Close() (err error) {
	return nil
}

var _ = Describe("Destinations create handler", func() {
	var (
		expectedResponseBody  []byte
		request               *http.Request
		handler               *handlers.DestinationsCreate
		resp                  *httptest.ResponseRecorder
		fakeMetricsSender     *storeFakes.MetricsSender
		fakeStore             *fakes.EgressDestinationStoreCreator
		fakeMarshaller        *fakes.EgressDestinationMarshaller
		fakePolicyGuard       *fakes.PolicyGuard
		logger                *lagertest.TestLogger
		createdDestinations   []store.EgressDestination
		requestedDestinations []store.EgressDestination
		token                 uaa_client.CheckTokenResponse
	)

	BeforeEach(func() {
		expectedResponseBody = []byte("some-response")

		requestBody := `{
					"destinations": [
						{  "name": "my service",
						    "description": "my service is a great service",	
						    "ips": [{"start": "7211.30.35.9", "end": "72.30.35.9"}],
						     "ports": [{"start": 8080, "end": 8080}],
						     "protocol":"tcp"
						},
						{  "name": "cloud infra",
						    "description": "this is where my apps go",
						    "ips": [{"start": "7211.30.35.9", "end": "72.30.35.9"}],
						     "ports": [{"start": 8080, "end": 8080}],
						     "protocol":"tcp"
						}
					]
				}`

		var err error
		request, err = http.NewRequest("POST", "/networking/v1/external/destinations", bytes.NewBuffer([]byte(requestBody)))
		Expect(err).NotTo(HaveOccurred())

		createdDestinations = []store.EgressDestination{
			{GUID: "created-one"},
			{GUID: "created-two"},
		}

		fakeStore = &fakes.EgressDestinationStoreCreator{}
		fakeStore.CreateReturns(createdDestinations, nil)

		fakeMarshaller = &fakes.EgressDestinationMarshaller{}
		fakeMarshaller.AsBytesReturns(expectedResponseBody, nil)

		fakePolicyGuard = &fakes.PolicyGuard{}
		fakePolicyGuard.IsNetworkAdminReturns(true)

		requestedDestinations = []store.EgressDestination{
			{GUID: "req-one"},
			{GUID: "req-two"},
		}
		fakeMarshaller.AsEgressDestinationsReturns(requestedDestinations, nil)

		logger = lagertest.NewTestLogger("test")

		fakeMetricsSender = &storeFakes.MetricsSender{}

		errorResponse := &httperror.ErrorResponse{
			MetricsSender: fakeMetricsSender,
		}

		handler = &handlers.DestinationsCreate{
			ErrorResponse:           errorResponse,
			EgressDestinationStore:  fakeStore,
			EgressDestinationMapper: fakeMarshaller,
			PolicyGuard:             fakePolicyGuard,
			Logger:                  logger,
		}
		resp = httptest.NewRecorder()

		token = uaa_client.CheckTokenResponse{
			Scope:    []string{"some-scope", "network.admin"},
			UserID:   "some-user-id",
			UserName: "some-user",
		}
	})

	It("creates destinations", func() {
		MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)

		Expect(fakePolicyGuard.IsNetworkAdminCallCount()).To(Equal(1))
		passedToken := fakePolicyGuard.IsNetworkAdminArgsForCall(0)
		Expect(passedToken).To(Equal(token))

		Expect(fakeStore.CreateCallCount()).To(Equal(1))
		Expect(fakeStore.CreateArgsForCall(0)).To(Equal(requestedDestinations))
		Expect(fakeMarshaller.AsBytesCallCount()).To(Equal(1))
		Expect(fakeMarshaller.AsBytesArgsForCall(0)).To(Equal(createdDestinations))
		Expect(resp.Code).To(Equal(http.StatusCreated))
		Expect(resp.Body.Bytes()).To(Equal(expectedResponseBody))
	})

	It("returns an error request body can't be read", func() {
		request.Body = &failingReader{}
		MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)

		Expect(resp.Code).To(Equal(http.StatusInternalServerError))
		Expect(resp.Body.Bytes()).To(MatchJSON(`{"error": "error reading request"}`))
	})

	It("returns an error when the store returns an error", func() {
		fakeStore.CreateReturns(nil, errors.New("can't create"))
		MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)
		Expect(resp.Code).To(Equal(http.StatusInternalServerError))
		Expect(resp.Body.Bytes()).To(MatchJSON(`{"error": "error creating egress destinations"}`))
	})

	It("returns an error when the mapper returns an error", func() {
		fakeMarshaller.AsEgressDestinationsReturns(nil, errors.New("whoa"))
		MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)
		Expect(resp.Code).To(Equal(http.StatusInternalServerError))
		Expect(resp.Body.Bytes()).To(MatchJSON(`{"error": "error parsing egress destinations"}`))
	})

	It("returns an error when the marshalling created destinations", func() {
		fakeMarshaller.AsBytesReturns(nil, errors.New("can't serialize"))
		MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)
		Expect(resp.Code).To(Equal(http.StatusInternalServerError))
		Expect(resp.Body.Bytes()).To(MatchJSON(`{"error": "error serializing egress destinations"}`))
	})

	Context("when the user is not network admin", func() {
		BeforeEach(func() {
			fakePolicyGuard.IsNetworkAdminReturns(false)
		})

		It("returns an error", func() {
			MakeRequestWithLoggerAndAuth(handler.ServeHTTP, resp, request, logger, token)
			Expect(resp.Code).To(Equal(http.StatusForbidden))
			Expect(resp.Body.Bytes()).To(MatchJSON(`{"error": "not authorized: creating egress destinations failed"}`))
		})
	})
})
