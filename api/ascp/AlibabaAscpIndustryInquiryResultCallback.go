package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryInquiryResultCallback 送货入户并安装服务商询价结果返回
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
func AlibabaAscpIndustryInquiryResultCallback(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaAscpIndustryInquiryResultCallbackAPIRequest, resp *ascp.AlibabaAscpIndustryInquiryResultCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
