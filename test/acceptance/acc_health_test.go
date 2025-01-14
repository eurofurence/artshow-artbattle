package acceptance

import (
	"github.com/eurofurence/artshow-artbattle/docs"
	"github.com/eurofurence/artshow-artbattle/internal/apimodel"
	"net/http"
	"testing"
)

// ----------------------------------------
// acceptance tests for the health endpoint
// ----------------------------------------

func TestHealthEndpoint(t *testing.T) {
	tstSetup(t)
	defer tstShutdown()

	docs.Given("given an anonymous user")

	docs.When("when they access the health endpoint")
	response := tstPerformGet("/", tstNoToken())

	docs.Then("then the operation is successful")
	tstRequireSuccessResponse(t, response, http.StatusOK, &apimodel.Health{Status: "OK"})
}
