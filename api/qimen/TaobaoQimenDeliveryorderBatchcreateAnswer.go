package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryorderbatchcreateanswer 发货单创建结果通知接口(批量)
// taobao.qimen.deliveryorder.batchcreate.answer
//
// WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
func Taobaoqimendeliveryorderbatchcreateanswer(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryorderbatchcreateanswerAPIRequest, session string) (*qimen.TaobaoqimendeliveryorderbatchcreateanswerAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryorderbatchcreateanswerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
