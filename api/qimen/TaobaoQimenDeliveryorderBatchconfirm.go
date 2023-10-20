package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryorderbatchconfirm 发货单确认接口
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
func Taobaoqimendeliveryorderbatchconfirm(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryorderbatchconfirmAPIRequest, session string) (*qimen.TaobaoqimendeliveryorderbatchconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryorderbatchconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
