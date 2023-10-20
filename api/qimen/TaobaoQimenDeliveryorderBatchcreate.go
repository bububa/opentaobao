package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryorderbatchcreate 发货单创建批量接口
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
func Taobaoqimendeliveryorderbatchcreate(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryorderbatchcreateAPIRequest, session string) (*qimen.TaobaoqimendeliveryorderbatchcreateAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryorderbatchcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
