package router

import (
	"log"
	"net/http"
	"regexp"
)

type routeRule struct {
	name    string
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	rules           []*routeRule
	notFoundHandler http.Handler
}

func NewRouter(notFoundHandler http.Handler) *Router {
	return &Router{
		rules:           make([]*routeRule, 0),
		notFoundHandler: notFoundHandler,
	}
}

func (rtr *Router) AddRule(name string, method, pattern string, handler http.HandlerFunc) {
	newRule := &routeRule{
		name:    name,
		method:  method,
		pattern: regexp.MustCompile(pattern),
		handler: handler,
	}
	rtr.rules = append(rtr.rules, newRule)
}

func (rtr *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	for _, rule := range rtr.rules {
		if rule.method != method {
			continue
		}
		if !rule.pattern.MatchString(path) {
			continue
		}
		log.Printf("found handler: %q", rule.name)
		handler := rule.handler
		handler.ServeHTTP(resp, req)
		return
	}
	// no rule for request
	rtr.notFoundHandler.ServeHTTP(resp, req)
}
