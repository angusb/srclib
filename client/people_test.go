package client

import (
	"net/http"
	"reflect"
	"testing"

	"sourcegraph.com/sourcegraph/api_router"
	"sourcegraph.com/sourcegraph/srcgraph/person"
)

func TestPeopleService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &person.User{UID: 1}

	var called bool
	mux.HandleFunc(urlPath(t, api_router.Person, map[string]string{"PersonSpec": "a"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	person_, _, err := client.People.Get(PersonSpec{LoginOrEmail: "a"})
	if err != nil {
		t.Errorf("People.Get returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(person_, want) {
		t.Errorf("People.Get returned %+v, want %+v", person_, want)
	}
}

func TestPeopleService_List(t *testing.T) {
	setup()
	defer teardown()

	want := []*person.User{{UID: 1}}

	var called bool
	mux.HandleFunc(urlPath(t, api_router.People, nil), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"Query":     "q",
			"Sort":      "name",
			"Direction": "asc",
			"PerPage":   "1",
			"Page":      "2",
		})

		writeJSON(w, want)
	})

	people, _, err := client.People.List(&PersonListOptions{
		Query:       "q",
		Sort:        "name",
		Direction:   "asc",
		ListOptions: ListOptions{PerPage: 1, Page: 2},
	})
	if err != nil {
		t.Errorf("People.List returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(people, want) {
		t.Errorf("People.List returned %+v, want %+v", people, want)
	}
}

func TestPeopleService_ListAuthors(t *testing.T) {
	setup()
	defer teardown()

	want := []*AugmentedPersonRef{{User: &person.User{UID: 1}}}

	var called bool
	mux.HandleFunc(urlPath(t, api_router.PersonAuthors, map[string]string{"PersonSpec": "a"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	authors, _, err := client.People.ListAuthors(PersonSpec{LoginOrEmail: "a"}, nil)
	if err != nil {
		t.Errorf("People.ListAuthors returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(authors, want) {
		t.Errorf("People.ListAuthors returned %+v, want %+v", authors, want)
	}
}

func TestPeopleService_ListClients(t *testing.T) {
	setup()
	defer teardown()

	want := []*AugmentedPersonRef{{User: &person.User{UID: 1}}}

	var called bool
	mux.HandleFunc(urlPath(t, api_router.PersonClients, map[string]string{"PersonSpec": "a"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	clients, _, err := client.People.ListClients(PersonSpec{LoginOrEmail: "a"}, nil)
	if err != nil {
		t.Errorf("People.ListClients returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(clients, want) {
		t.Errorf("People.ListClients returned %+v, want %+v", clients, want)
	}
}