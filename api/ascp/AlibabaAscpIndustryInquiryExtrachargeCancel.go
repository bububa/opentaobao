package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryInquiryExtrachargeCancel 送货入户并安装服务商取消增加费用
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
func AlibabaAscpIndustryInquiryExtrachargeCancel(clt *core.SDKClient, req *ascp.AlibabaAscpIndustryInquiryExtrachargeCancelAPIRequest, session string) (*ascp.AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse, error) {
	var resp ascp.AlibabaAscpIndustryInquiryExtrachargeCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
