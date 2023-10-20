package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabaascpindustryinquiryextrachargecancel 送货入户并安装服务商取消增加费用
// alibaba.ascp.industry.inquiry.extracharge.cancel
//
// 送货入户并安装服务商取消增加费用
func Alibabaascpindustryinquiryextrachargecancel(clt *core.SDKClient, req *ascp.AlibabaascpindustryinquiryextrachargecancelAPIRequest, session string) (*ascp.AlibabaascpindustryinquiryextrachargecancelAPIResponse, error) {
	var resp ascp.AlibabaascpindustryinquiryextrachargecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
