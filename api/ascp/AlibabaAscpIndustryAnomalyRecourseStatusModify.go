package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryAnomalyRecourseStatusModify 送货入户并安装投诉工单状态变更
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
func AlibabaAscpIndustryAnomalyRecourseStatusModify(clt *core.SDKClient, req *ascp.AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIRequest, resp *ascp.AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
