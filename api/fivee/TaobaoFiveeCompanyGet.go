package fivee

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeCompanyGet 查询商信息
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
func TaobaoFiveeCompanyGet(ctx context.Context, clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyGetAPIRequest, resp *fivee.TaobaoFiveeCompanyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
