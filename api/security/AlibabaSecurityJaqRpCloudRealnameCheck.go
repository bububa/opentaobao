package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudrealnamecheck 验证姓名和证件号
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
func Alibabasecurityjaqrpcloudrealnamecheck(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudrealnamecheckAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudrealnamecheckAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudrealnamecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
