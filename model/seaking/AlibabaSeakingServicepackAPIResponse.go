package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingservicepackAPIResponse 获取海王用户权限包 API返回值
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
type AlibabaseakingservicepackAPIResponse struct {
	model.CommonResponse
	AlibabaseakingservicepackAPIResponseModel
}

// AlibabaseakingservicepackAPIResponseModel is 获取海王用户权限包 成功返回结果
type AlibabaseakingservicepackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_servicepack_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权限包列表
	ServicePackList []AlibabaseakingservicepackResult `json:"service_pack_list,omitempty" xml:"service_pack_list>alibabaseakingservicepack_result,omitempty"`
}
