package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaInteractOrderCheckuserimei 金融购机验证设备号
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
func AlibabaInteractOrderCheckuserimei(clt *core.SDKClient, req *alicom.AlibabaInteractOrderCheckuserimeiAPIRequest, session string) (*alicom.AlibabaInteractOrderCheckuserimeiAPIResponse, error) {
	var resp alicom.AlibabaInteractOrderCheckuserimeiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
