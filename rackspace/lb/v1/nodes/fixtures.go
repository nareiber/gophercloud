package nodes

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func _rootURL(lbID int) string {
	return "/loadbalancers/" + strconv.Itoa(lbID) + "/nodes"
}

func mockListResponse(t *testing.T, lbID int) {
	th.Mux.HandleFunc(_rootURL(lbID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
  "nodes": [
    {
      "id": 410,
      "address": "10.1.1.1",
      "port": 80,
      "condition": "ENABLED",
      "status": "ONLINE",
      "weight": 3,
      "type": "PRIMARY"
    },
    {
      "id": 411,
      "address": "10.1.1.2",
      "port": 80,
      "condition": "ENABLED",
      "status": "ONLINE",
      "weight": 8,
      "type": "SECONDARY"
    }
  ]
}
  `)
	})
}

func mockCreateResponse(t *testing.T, lbID int) {
	th.Mux.HandleFunc(_rootURL(lbID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.TestJSONRequest(t, r, `
{
  "nodes": [
    {
      "address": "10.2.2.3",
      "port": 80,
      "condition": "ENABLED",
      "type": "PRIMARY"
    },
    {
      "address": "10.2.2.4",
      "port": 81,
      "condition": "ENABLED",
      "type": "SECONDARY"
    }
  ]
}
    `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
  "nodes": [
    {
      "address": "10.2.2.3",
      "id": 185,
      "port": 80,
      "status": "ONLINE",
      "condition": "ENABLED",
      "weight": 1,
      "type": "PRIMARY"
    },
    {
      "address": "10.2.2.4",
      "id": 186,
      "port": 81,
      "status": "ONLINE",
      "condition": "ENABLED",
      "weight": 1,
      "type": "SECONDARY"
    }
  ]
}
  `)
	})
}

func mockBatchDeleteLBResponse(t *testing.T, ids []int) {
	th.Mux.HandleFunc("/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		r.ParseForm()

		for k, v := range ids {
			fids := r.Form["id"]
			th.AssertEquals(t, strconv.Itoa(v), fids[k])
		}

		w.WriteHeader(http.StatusAccepted)
	})
}

func mockDeleteLBResponse(t *testing.T, id int) {
	th.Mux.HandleFunc("/loadbalancers/"+strconv.Itoa(id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func mockGetResponse(t *testing.T, id int) {
	th.Mux.HandleFunc("/loadbalancers/"+strconv.Itoa(id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
  "node": {
    "id": 410,
    "address": "10.1.1.1",
    "port": 80,
    "condition": "ENABLED",
    "status": "ONLINE",
    "weight": 12,
    "type": "PRIMARY"
  }
}
  `)
	})
}

func mockUpdateResponse(t *testing.T, id int) {
	th.Mux.HandleFunc("/loadbalancers/"+strconv.Itoa(id), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.TestJSONRequest(t, r, `
{
  "node": {
    "condition": "ENABLED",
    "weight": 59
  }
}
    `)

		w.WriteHeader(http.StatusOK)
	})
}

func mockListEventsResponse(t *testing.T) {
	th.Mux.HandleFunc("/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
  "nodeServiceEvents": [
    {
      "detailedMessage": "Node is ok",
      "nodeId": 373,
      "id": 7,
      "type": "UPDATE_NODE",
      "description": "Node '373' status changed to 'ONLINE' for load balancer '323'",
      "category": "UPDATE",
      "severity": "INFO",
      "relativeUri": "/406271/loadbalancers/323/nodes/373/events",
      "accountId": 406271,
      "loadbalancerId": 323,
      "title": "Node Status Updated",
      "author": "Rackspace Cloud",
      "created": "10-30-2012 10:18:23"
    },
    {
      "detailedMessage": "Invalid HTTP response received; premature end of headers",
      "nodeId": 373,
      "id": 8,
      "type": "UPDATE_NODE",
      "description": "Node '373' status changed to 'OFFLINE' for load balancer '323'",
      "category": "UPDATE",
      "severity": "INFO",
      "relativeUri": "/406271/loadbalancers/323/nodes/373/events",
      "accountId": 406271,
      "loadbalancerId": 323,
      "title": "Node Status Updated",
      "author": "Rackspace Cloud",
      "created": "10-30-2012 11:22:25"
    }
  ]
}
`)
	})
}
