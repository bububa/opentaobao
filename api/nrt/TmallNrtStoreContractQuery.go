package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreContractQuery 摊位合同查询接口
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
func TmallNrtStoreContractQuery(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtStoreContractQueryAPIRequest, resp *nrt.TmallNrtStoreContractQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
