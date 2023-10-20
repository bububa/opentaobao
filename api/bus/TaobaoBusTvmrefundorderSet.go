package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTvmrefundorderSet 线下自助机逆向退款接口
// taobao.bus.tvmrefundorder.set
//
// 汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
func TaobaoBusTvmrefundorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmrefundorderSetAPIRequest, resp *bus.TaobaoBusTvmrefundorderSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
