package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizApAdd 新增价签通讯AP设备
// taobao.uscesl.biz.ap.add
//
// 根据门店和ap的MAC地址新增
func TaobaoUsceslBizApAdd(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApAddAPIRequest, resp *uscesl.TaobaoUsceslBizApAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
