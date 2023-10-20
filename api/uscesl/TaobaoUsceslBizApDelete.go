package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizApDelete 删除价签AP设备
// taobao.uscesl.biz.ap.delete
//
// 删除价签AP设备
func TaobaoUsceslBizApDelete(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApDeleteAPIRequest, resp *uscesl.TaobaoUsceslBizApDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
