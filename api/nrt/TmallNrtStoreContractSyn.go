package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreContractSyn 喵零合同同步
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
func TmallNrtStoreContractSyn(clt *core.SDKClient, req *nrt.TmallNrtStoreContractSynAPIRequest, session string) (*nrt.TmallNrtStoreContractSynAPIResponse, error) {
	var resp nrt.TmallNrtStoreContractSynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
