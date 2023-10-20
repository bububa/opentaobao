package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketTasksGet 任务列表获取接口
// taobao.vmarket.eticket.tasks.get
//
// 外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
func TaobaoVmarketEticketTasksGet(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketTasksGetAPIRequest, resp *eticket.TaobaoVmarketEticketTasksGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
