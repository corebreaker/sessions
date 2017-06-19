# Session

[![Build Status](https://travis-ci.org/corebreker/sessions.svg?branch=master)](https://travis-ci.org/corebreker/sessions)
[![GoDoc](https://godoc.org/github.com/corebreker/sessions?status.svg)](https://godoc.org/github.com/corebreker/sessions)
[![Go Report Card](https://goreportcard.com/badge/github.com/corebreker/sessions)](https://goreportcard.com/report/github.com/corebreker/sessions)
[![codecov](https://codecov.io/gh/corebreker/sessions/branch/master/graph/badge.svg)](https://codecov.io/gh/corebreker/sessions)

The [Go](https://golang.org/) standard library includes a nice [http server](https://golang.org/pkg/net/http/), but unfortunately it lacks a very basic and important feature: _HTTP session management_.

This package provides an easy-to-use, extensible and secure session implementation and management. Package documentation can be found and godoc.org:

[https://godoc.org/github.com/corebreker/sessions](https://godoc.org/github.com/corebreker/sessions)

This is "just" an session data implementation and management, you can use it as-is, or with any existing Go web toolkits and frameworks. But you can use it in a personnal client/server that need to associate datas for a logged user.

## Overview

There are 3 key _players_ in the package:

- **`Session`** is the (HTTP) session interface. We can use it to store and retrieve constant and variable attributes from it.
- **`Store`** is a session store interface which is responsible to store sessions and make them retrievable by their IDs at the server side.
- **`Manager`** is a session manager interface which is responsible to acquire a `Session` from an (incoming) HTTP request, and to add a `Session` to an HTTP response to let the client know about the session. A `Manager` has a backing `Store` which is responsible to manage `Session` values at server side.

_Players_ of this package are represented by interfaces, and various implementations are provided for all these players.
You are not bound by the provided implementations, feel free to provide your own implementations for any of the players.

## Usage

Usage can't be simpler than this. To get the current session associated with the [http.Request](https://golang.org/pkg/net/http/#Request):

    sess := session.Get(r)
    if sess == nil {
        // No session (yet)
    } else {
        // We have a session, use it
    }

To create a new session (e.g. on a successful login) and add it to an [http.ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) (to let the client know about the session):

    sess := session.NewSession()
    session.Add(sess, w)

Let's see a more advanced session creation: let's provide a constant attribute (for the lifetime of the session) and an initial, variable attribute:

    sess := session.NewSessionOptions(&session.SessOptions{
        CAttrs: map[string]interface{}{"UserName": userName},
        Attrs:  map[string]interface{}{"Count": 1},
    })

And to access these attributes and change value of `"Count"`:

    userName := sess.CAttr("UserName")
    count := sess.Attr("Count").(int) // Type assertion, you might wanna check if it succeeds
    sess.SetAttr("Count", count+1)    // Increment count

(Of course variable attributes can be added later on too with `Session.SetAttr()`, not just at session creation.)

To remove a session (e.g. on logout):

    session.Remove(sess, w)

Check out the [session demo application](https://github.com/corebreker/sessions/blob/master/session_demo/session_demo.go) which shows all these in action.
