package utils

import (
	"context"
	"testing"

	"github.com/asim/go-micro/plugins/registry/memory/v3"
	breaker "github.com/asim/go-micro/plugins/wrapper/breaker/gobreaker/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/selector"
	"github.com/sony/gobreaker"
)

func TestBreaker(t *testing.T) {
	// setup
	r := memory.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := client.NewClient(
		// set the selector
		client.Selector(s),
		// add the breaker wrapper
		client.Wrap(breaker.NewClientWrapper()),
	)

	req := c.NewRequest("test.service", "Test.Method", map[string]string{
		"foo": "bar",
	}, client.WithContentType("application/json"))

	var rsp map[string]any

	// Force to point of trip
	for i := 0; i < 6; i++ {
		c.Call(context.TODO(), req, rsp)
	}

	err := c.Call(context.TODO(), req, rsp)
	if err == nil {
		t.Error("Expecting tripped breaker, got nil error")
	}

	merr := err.(*errors.Error)
	if merr.Code != 502 {
		t.Errorf("Expecting tripped breaker, got %v", err)
	}
}

func TestCustomBreaker(t *testing.T) {
	// setup
	r := memory.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := client.NewClient(
		// set the selector
		client.Selector(s),
		// add the breaker wrapper
		client.Wrap(breaker.NewCustomClientWrapper(
			gobreaker.Settings{},
			breaker.BreakService,
		)),
	)

	req := c.NewRequest("test.service", "Test.Method", map[string]string{
		"foo": "bar",
	}, client.WithContentType("application/json"))

	var rsp map[string]any

	// Force to point of trip
	for i := 0; i < 6; i++ {
		c.Call(context.TODO(), req, rsp)
	}

	err := c.Call(context.TODO(), req, rsp)
	if err == nil {
		t.Error("Expecting tripped breaker, got nil error")
	}

	merr := err.(*errors.Error)
	if merr.Code != 502 {
		t.Errorf("Expecting tripped breaker, got %v", err)
	}
}
