package ticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// Alitripticketproductquery 【门票API2.0】门票商品查询接口
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
func Alitripticketproductquery(clt *core.SDKClient, req *ticket.AlitripticketproductqueryAPIRequest, session string) (*ticket.AlitripticketproductqueryAPIResponse, error) {
	var resp ticket.AlitripticketproductqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
