package example

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/memcache"
)

func TestFoo(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	it := &memcache.Item{
		Key:   "some-key",
		Value: []byte("some-value"),
	}
	err = memcache.Set(ctx, it)
	if err != nil {
		t.Fatalf("Set err: %v", err)
	}
	it, err = memcache.Get(ctx, "some-key")
	if err != nil {
		t.Fatalf("Get err: %v; want no error", err)
	}
	if g, w := string(it.Value), "some-value"; g != w {
		t.Errorf("retrieved Item.Value = %q, want %q", g, w)
	}
}

func TestServiceStatus(t *testing.T) {
	instance, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create the app engine context: %v", err)
	}
	defer instance.Close()

	req, err := instance.NewRequest("GET", "/service/status", nil)
	if err != nil {
		t.Fatalf("Failed to create GET request to /service/status: %v", err)
	}

	resp := httptest.NewRecorder()

	// use the original handler
	ServiceStatus(resp, req, nil)

	st := &serviceStatus{}

	err = json.Unmarshal(resp.Body.Bytes(), st)
	if err != nil {
		t.Errorf("Error on unmarshaling the response: ", err)
	}

	// ASSERT Equal: final comparison between the local mock and sent mock data
	if st.Message != "Micro Service EXAMPLE Status is OK" {
		t.Errorf("Error, different message!")
	}

	// ASSERT Equal: final comparison between the local mock and sent mock data
	if resp.Code != http.StatusOK {
		t.Errorf("Error, different HTTP STATUS!")
	}

}
