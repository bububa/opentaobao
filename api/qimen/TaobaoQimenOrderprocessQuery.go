package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenorderprocessquery 订单流水查询接口
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
func Taobaoqimenorderprocessquery(clt *core.SDKClient, req *qimen.TaobaoqimenorderprocessqueryAPIRequest, session string) (*qimen.TaobaoqimenorderprocessqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenorderprocessqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
