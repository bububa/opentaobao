package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqResourceFetch 获取资源文件
// alibaba.security.jaq.resource.fetch
//
// 在前向化验证流程中提供资源文件服务
func AlibabaSecurityJaqResourceFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqResourceFetchAPIRequest, resp *security.AlibabaSecurityJaqResourceFetchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
