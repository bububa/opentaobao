package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqResourceFetch 获取资源文件
// alibaba.security.jaq.resource.fetch
//
// 在前向化验证流程中提供资源文件服务
func AlibabaSecurityJaqResourceFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqResourceFetchAPIRequest, session string) (*security.AlibabaSecurityJaqResourceFetchAPIResponse, error) {
	var resp security.AlibabaSecurityJaqResourceFetchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
