package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryAnomalyRecourseStatusModify 送货入户并安装投诉工单状态变更
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
func AlibabaAscpIndustryAnomalyRecourseStatusModify(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest, resp *ascp.AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
