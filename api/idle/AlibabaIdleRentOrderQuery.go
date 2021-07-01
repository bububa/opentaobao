package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

/* AlibabaIdleRentOrderQuery
查询订单
alibaba.idle.rent.order.query

查询订单信息 */
func AlibabaIdleRentOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderQueryAPIRequest, session string) (*idle.AlibabaIdleRentOrderQueryAPIResponse, error) {
	var resp idle.AlibabaIdleRentOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
