package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponse 服务商工人客诉数据上传 API返回值
// alibaba.dchain.miaoshifu.customer.complaints.put
//
// 数字服务供应链平台提供给服务商上传工人客诉数据
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponse struct {
	model.CommonResponse
	AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponseModel
}

// AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponseModel is 服务商工人客诉数据上传 成功返回结果
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_miaoshifu_customer_complaints_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功，true：成功，false失败，当未false时，result_value为null
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
	// 错误码，异常返回码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 客诉记录唯一标识ID
	ResultValue int64 `json:"result_value,omitempty" xml:"result_value,omitempty"`
	// 错误信息
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
}
