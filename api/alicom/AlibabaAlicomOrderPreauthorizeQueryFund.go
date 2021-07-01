package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAlicomOrderPreauthorizeQueryFund
资金流水查询
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询 */
func AlibabaAlicomOrderPreauthorizeQueryFund(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest, session string) (*alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse, error) {
	var resp alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
