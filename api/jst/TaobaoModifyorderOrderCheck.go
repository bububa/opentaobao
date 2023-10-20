package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoModifyorderOrderCheck 自助改单服务发货订单校验
// taobao.modifyorder.order.check
//
// 自助改单服务发货后订单回传接口
func TaobaoModifyorderOrderCheck(clt *core.SDKClient, req *jst.TaobaoModifyorderOrderCheckAPIRequest, resp *jst.TaobaoModifyorderOrderCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
