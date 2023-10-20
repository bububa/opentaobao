package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectticketquery 根据商品id查询核销卷信息
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
func Alibabaalihousenewhomeprojectticketquery(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectticketqueryAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectticketqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectticketqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
