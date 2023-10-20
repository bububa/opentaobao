package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstorecontractsyn 喵零合同同步
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
func Tmallnrtstorecontractsyn(clt *core.SDKClient, req *nrt.TmallnrtstorecontractsynAPIRequest, session string) (*nrt.TmallnrtstorecontractsynAPIResponse, error) {
	var resp nrt.TmallnrtstorecontractsynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
