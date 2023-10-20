package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudstart 实人认证云开始认证
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
func Alibabasecurityjaqrpcloudstart(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudstartAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudstartAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudstartAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
