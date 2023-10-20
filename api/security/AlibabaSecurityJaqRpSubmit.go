package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpsubmit 聚安全实人认证提交认证接口
// alibaba.security.jaq.rp.submit
//
// 聚安全实人认证提交认证接口
func Alibabasecurityjaqrpsubmit(clt *core.SDKClient, req *security.AlibabasecurityjaqrpsubmitAPIRequest, session string) (*security.AlibabasecurityjaqrpsubmitAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
