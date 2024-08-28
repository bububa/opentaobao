package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqWsgriskdataReport 无线保镖SDK风控数据上报
// alibaba.security.jaq.wsgriskdata.report
//
// 无线保镖sdk根据用户的需要，上报数据到聚安全云端
func AlibabaSecurityJaqWsgriskdataReport(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqWsgriskdataReportAPIRequest, resp *security.AlibabaSecurityJaqWsgriskdataReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
