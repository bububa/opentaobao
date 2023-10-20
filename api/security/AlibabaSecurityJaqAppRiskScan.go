package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppRiskScan 应用风险扫描提交接口
// alibaba.security.jaq.app.risk.scan
//
// 提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
func AlibabaSecurityJaqAppRiskScan(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskScanAPIRequest, session string) (*security.AlibabaSecurityJaqAppRiskScanAPIResponse, error) {
	var resp security.AlibabaSecurityJaqAppRiskScanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
