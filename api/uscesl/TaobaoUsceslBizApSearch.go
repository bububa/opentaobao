package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizApSearch AP列表查询
// taobao.uscesl.biz.ap.search
//
// 查询当前门店下登记的AP列表
func TaobaoUsceslBizApSearch(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApSearchAPIRequest, resp *uscesl.TaobaoUsceslBizApSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
