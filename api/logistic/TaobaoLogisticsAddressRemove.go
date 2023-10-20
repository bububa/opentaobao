package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressRemove 删除卖家地址库
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
func TaobaoLogisticsAddressRemove(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressRemoveAPIRequest, resp *logistic.TaobaoLogisticsAddressRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
