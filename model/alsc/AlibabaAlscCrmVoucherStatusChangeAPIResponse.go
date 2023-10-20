package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmvoucherstatuschangeAPIResponse 优惠券状态更改 API返回值
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
type AlibabaalsccrmvoucherstatuschangeAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmvoucherstatuschangeAPIResponseModel
}

// AlibabaalsccrmvoucherstatuschangeAPIResponseModel is 优惠券状态更改 成功返回结果
type AlibabaalsccrmvoucherstatuschangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
