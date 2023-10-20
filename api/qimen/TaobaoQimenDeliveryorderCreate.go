package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryordercreate 发货单创建接口
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
func Taobaoqimendeliveryordercreate(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryordercreateAPIRequest, session string) (*qimen.TaobaoqimendeliveryordercreateAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
