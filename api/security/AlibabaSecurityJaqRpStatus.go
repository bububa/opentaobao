package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpstatus 聚安全实人认证查询状态接口
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
func Alibabasecurityjaqrpstatus(clt *core.SDKClient, req *security.AlibabasecurityjaqrpstatusAPIRequest, session string) (*security.AlibabasecurityjaqrpstatusAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
