package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppRiskdetailbatchGet 应用风险详细信息批量查询接口
// alibaba.security.jaq.app.riskdetailbatch.get
//
// 用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
func AlibabaSecurityJaqAppRiskdetailbatchGet(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest, session string) (*security.AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse, error) {
	var resp security.AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
