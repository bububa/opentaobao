package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunospubadmincommonoperationAPIResponse 内部迎客松通用服务 API返回值
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
type YunospubadmincommonoperationAPIResponse struct {
	model.CommonResponse
	YunospubadmincommonoperationAPIResponseModel
}

// YunospubadmincommonoperationAPIResponseModel is 内部迎客松通用服务 成功返回结果
type YunospubadmincommonoperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_pubadmin_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
