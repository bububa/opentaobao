package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomvtdistributesendcode 通信业务外放发送验证码
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
func Alibabaalicomvtdistributesendcode(clt *core.SDKClient, req *alicom.AlibabaalicomvtdistributesendcodeAPIRequest, session string) (*alicom.AlibabaalicomvtdistributesendcodeAPIResponse, error) {
	var resp alicom.AlibabaalicomvtdistributesendcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
