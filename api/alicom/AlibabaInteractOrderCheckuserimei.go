package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabainteractordercheckuserimei 金融购机验证设备号
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
func Alibabainteractordercheckuserimei(clt *core.SDKClient, req *alicom.AlibabainteractordercheckuserimeiAPIRequest, session string) (*alicom.AlibabainteractordercheckuserimeiAPIResponse, error) {
	var resp alicom.AlibabainteractordercheckuserimeiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
