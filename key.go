package sessions

type Key interface {
    GetId() SessionID
    SetId(sid SessionID)
}
