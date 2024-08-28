package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoModifyskuQueryStatus 查询商家是否开通自助修改商品信息服务
// taobao.modifysku.query.status
//
// 查询商家是否开通自助修改商品信息服务
// 1. 没有授权
// 2. 已与当前appkey签约
// 3. 没有签约
// 4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
func TaobaoModifyskuQueryStatus(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoModifyskuQueryStatusAPIRequest, resp *jst.TaobaoModifyskuQueryStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
