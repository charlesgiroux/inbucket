package rest

import "github.com/gorilla/mux"
import "github.com/jhillyerd/inbucket/pkg/server/web"

// NewRouter returns a router configured for the REST interface.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// API v1
	r.Path("/api/v1/mailbox/{name}").Handler(
		web.Handler(MailboxListV1)).Name("MailboxListV1").Methods("GET")
	r.Path("/api/v1/mailbox/{name}").Handler(
		web.Handler(MailboxPurgeV1)).Name("MailboxPurgeV1").Methods("DELETE")
	r.Path("/api/v1/mailbox/{name}/{id}").Handler(
		web.Handler(MailboxShowV1)).Name("MailboxShowV1").Methods("GET")
	r.Path("/api/v1/mailbox/{name}/{id}").Handler(
		web.Handler(MailboxMarkSeenV1)).Name("MailboxMarkSeenV1").Methods("PATCH")
	r.Path("/api/v1/mailbox/{name}/{id}").Handler(
		web.Handler(MailboxDeleteV1)).Name("MailboxDeleteV1").Methods("DELETE")
	r.Path("/api/v1/mailbox/{name}/{id}/source").Handler(
		web.Handler(MailboxSourceV1)).Name("MailboxSourceV1").Methods("GET")
	r.Path("/api/v1/monitor/messages").Handler(
		web.Handler(MonitorAllMessagesV1)).Name("MonitorAllMessagesV1").Methods("GET")
	r.Path("/api/v1/monitor/messages/{name}").Handler(
		web.Handler(MonitorMailboxMessagesV1)).Name("MonitorMailboxMessagesV1").Methods("GET")
	return r
}
