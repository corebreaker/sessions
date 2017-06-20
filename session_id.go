package sessions

import (
    "crypto/rand"
    "encoding/base64"
    "io"
)

const NO_SESSION SessionID = ""

type SessionID string

func (sid SessionID) HasSession() bool {
    return sid != NO_SESSION
}

// genID generates a secure, random session id using the crypto/rand package.
func genID(length int) SessionID {
    r := make([]byte, base64.URLEncoding.DecodedLen(length+3))
    io.ReadFull(rand.Reader, r)

    res := base64.URLEncoding.EncodeToString(r)

    return res[:length]
}
