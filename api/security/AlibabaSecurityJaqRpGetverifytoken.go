package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpgetverifytoken 聚安全实人认证获取认证会话token
// alibaba.security.jaq.rp.getverifytoken
//
// 聚安全实人认证获取认证会话token
func Alibabasecurityjaqrpgetverifytoken(clt *core.SDKClient, req *security.AlibabasecurityjaqrpgetverifytokenAPIRequest, session string) (*security.AlibabasecurityjaqrpgetverifytokenAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpgetverifytokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
