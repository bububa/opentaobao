package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryInquiryExtrachargeCancel 送货入户并安装服务商取消增加费用
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
func AlibabaAscpIndustryInquiryExtrachargeCancel(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest, resp *ascp.AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
