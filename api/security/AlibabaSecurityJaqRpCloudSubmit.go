package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqrpcloudsubmit 实人认证云服务提交接口
// alibaba.security.jaq.rp.cloud.submit
//
// 聚安全实人认证提交认证接口
func Alibabasecurityjaqrpcloudsubmit(clt *core.SDKClient, req *security.AlibabasecurityjaqrpcloudsubmitAPIRequest, session string) (*security.AlibabasecurityjaqrpcloudsubmitAPIResponse, error) {
	var resp security.AlibabasecurityjaqrpcloudsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
