package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoVasSubscSearch 订购记录导出
// taobao.vas.subsc.search
//
// 用于ISV查询自己名下的应用及收费项目的订购记录
func TaobaoVasSubscSearch(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoVasSubscSearchAPIRequest, resp *servicecenter.TaobaoVasSubscSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
