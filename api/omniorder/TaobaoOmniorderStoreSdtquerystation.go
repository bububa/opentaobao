package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSdtquerystation 速店通查询站点信息
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
func TaobaoOmniorderStoreSdtquerystation(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtquerystationAPIRequest, resp *omniorder.TaobaoOmniorderStoreSdtquerystationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
