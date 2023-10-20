package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreContractQuery 摊位合同查询接口
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
func TmallNrtStoreContractQuery(clt *core.SDKClient, req *nrt.TmallNrtStoreContractQueryAPIRequest, resp *nrt.TmallNrtStoreContractQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
