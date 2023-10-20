package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressModify 卖家地址库修改
// taobao.logistics.address.modify
//
// 卖家地址库修改
func TaobaoLogisticsAddressModify(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressModifyAPIRequest, resp *logistic.TaobaoLogisticsAddressModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
