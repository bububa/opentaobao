package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimendeliveryorderquery 发货单查询接口
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
func Taobaoqimendeliveryorderquery(clt *core.SDKClient, req *qimen.TaobaoqimendeliveryorderqueryAPIRequest, session string) (*qimen.TaobaoqimendeliveryorderqueryAPIResponse, error) {
	var resp qimen.TaobaoqimendeliveryorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
