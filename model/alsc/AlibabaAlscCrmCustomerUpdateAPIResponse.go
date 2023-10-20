package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomerupdateAPIResponse 更新顾客信息 API返回值
// alibaba.alsc.crm.customer.update
//
// 更新顾客信息
type AlibabaalsccrmcustomerupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmcustomerupdateAPIResponseModel
}

// AlibabaalsccrmcustomerupdateAPIResponseModel is 更新顾客信息 成功返回结果
type AlibabaalsccrmcustomerupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
