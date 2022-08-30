package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreContractQuery 摊位合同查询接口
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
func TmallNrtStoreContractQuery(clt *core.SDKClient, req *nrt.TmallNrtStoreContractQueryAPIRequest, session string) (*nrt.TmallNrtStoreContractQueryAPIResponse, error) {
	var resp nrt.TmallNrtStoreContractQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
