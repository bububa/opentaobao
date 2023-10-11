package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryInquiryResultCallback 送货入户并安装服务商询价结果返回
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
func AlibabaAscpIndustryInquiryResultCallback(clt *core.SDKClient, req *ascp.AlibabaAscpIndustryInquiryResultCallbackAPIRequest, session string) (*ascp.AlibabaAscpIndustryInquiryResultCallbackAPIResponse, error) {
	var resp ascp.AlibabaAscpIndustryInquiryResultCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
