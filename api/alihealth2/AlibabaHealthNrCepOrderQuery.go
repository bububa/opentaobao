package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaHealthNrCepOrderQuery
订单详情查询接口
alibaba.health.nr.cep.order.query

订单详情查询接口 */
func AlibabaHealthNrCepOrderQuery(clt *core.SDKClient, req *alihealth2.AlibabaHealthNrCepOrderQueryAPIRequest, session string) (*alihealth2.AlibabaHealthNrCepOrderQueryAPIResponse, error) {
	var resp alihealth2.AlibabaHealthNrCepOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
