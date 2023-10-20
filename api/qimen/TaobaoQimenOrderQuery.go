package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenorderquery 根据收件人信息查询交易单号接口
// taobao.qimen.order.query
//
// WMS 调用该接口，根据收件人信息查询平台交易订单号。
func Taobaoqimenorderquery(clt *core.SDKClient, req *qimen.TaobaoqimenorderqueryAPIRequest, session string) (*qimen.TaobaoqimenorderqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
