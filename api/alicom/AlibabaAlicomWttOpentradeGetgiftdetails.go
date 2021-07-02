package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomWttOpentradeGetgiftdetails 存送业务查询奖品信息
// alibaba.alicom.wtt.opentrade.getgiftdetails
//
// 话费宝充值送查询奖品信息
func AlibabaAlicomWttOpentradeGetgiftdetails(clt *core.SDKClient, req *alicom.AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest, session string) (*alicom.AlibabaAlicomWttOpentradeGetgiftdetailsAPIResponse, error) {
	var resp alicom.AlibabaAlicomWttOpentradeGetgiftdetailsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
