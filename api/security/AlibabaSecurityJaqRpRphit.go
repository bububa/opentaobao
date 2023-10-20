package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrprphit 聚安全-实人认证日志打点接口
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
func Alibabasecurityjaqrprphit(clt *core.SDKClient, req *security.AlibabasecurityjaqrprphitAPIRequest, session string) (*security.AlibabasecurityjaqrprphitAPIResponse, error) {
	var resp security.AlibabasecurityjaqrprphitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
