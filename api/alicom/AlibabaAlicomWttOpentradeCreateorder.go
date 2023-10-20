package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomwttopentradecreateorder 充值送活动下单接口
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
func Alibabaalicomwttopentradecreateorder(clt *core.SDKClient, req *alicom.AlibabaalicomwttopentradecreateorderAPIRequest, session string) (*alicom.AlibabaalicomwttopentradecreateorderAPIResponse, error) {
	var resp alicom.AlibabaalicomwttopentradecreateorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
