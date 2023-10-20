package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryorderconfirm 发货单确认接口
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
func Taobaoqimendeliveryorderconfirm(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryorderconfirmAPIRequest, session string) (*qimen.TaobaoqimendeliveryorderconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
