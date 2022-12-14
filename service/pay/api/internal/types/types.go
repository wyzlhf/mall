// Code generated by goctl. DO NOT EDIT.
package types

type CreateReuest struct {
	Uid    uint64 `json:"uid"`
	Oid    uint64 `json:"oid"`
	Amount uint64 `json:"amount"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type DetailRequest struct {
	Id uint64 `json:"id"`
}

type DetailResponse struct {
	Id     uint64 `json:"id"`
	Uid    uint64 `json:"uid"`
	Oid    uint64 `json:"oid"`
	Amount uint64 `json:"amount"`
	Source uint64 `json:"source"`
	Status uint64 `json:"status"`
}

type CallbackRequest struct {
	Id     uint64 `json:"id"`
	Uid    uint64 `json:"uid"`
	Oid    uint64 `json:"oid"`
	Amount uint64 `json:"amount"`
	Source uint64 `json:"source"`
	Status uint64 `json:"status"`
}

type CallbackResponse struct {
}
