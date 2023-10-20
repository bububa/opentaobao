package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvmboscommonoperationAPIResponse 应用中心通用服务接口 API返回值
// yunos.tvmbos.common.operation
//
// 应用中心相关接口的代理
type YunostvmboscommonoperationAPIResponse struct {
	model.CommonResponse
	YunostvmboscommonoperationAPIResponseModel
}

// YunostvmboscommonoperationAPIResponseModel is 应用中心通用服务接口 成功返回结果
type YunostvmboscommonoperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvmbos_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
