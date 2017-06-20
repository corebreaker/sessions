/*

Session Store interface.

*/

package sessions

// Store is a session store interface.
// A session store is responsible to store sessions and make them retrievable by their IDs at the server side.
type Store interface {
    // Get returns the session specified by its id.
    // The returned session will have an updated access time (set to the current time).
    // nil is returned if this store does not contain a session with the specified id.
    Get(id SessionID) Session

    // New creates and adds a new session to the store.
    New() Session

    // Add adds a session to the store.
    Add(sess Session)

    // Remove removes a session from the store.
    Remove(sess Session)

    // Close closes the session store, releasing any resources that were allocated.
    Close()
}
