package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCreateDo top创建DO/RT接口
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
func AlibabaPurCreateDo(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurCreateDoAPIRequest, resp *pur.AlibabaPurCreateDoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
