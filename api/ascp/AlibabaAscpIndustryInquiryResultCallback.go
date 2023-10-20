package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabaascpindustryinquiryresultcallback 送货入户并安装服务商询价结果返回
// alibaba.ascp.industry.inquiry.result.callback
//
// 送货入户并安装服务商询价结果返回
func Alibabaascpindustryinquiryresultcallback(clt *core.SDKClient, req *ascp.AlibabaascpindustryinquiryresultcallbackAPIRequest, session string) (*ascp.AlibabaascpindustryinquiryresultcallbackAPIResponse, error) {
	var resp ascp.AlibabaascpindustryinquiryresultcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
