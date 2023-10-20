package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponse 空间单元列表带业务属性实例 API返回值
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
type AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponse struct {
	model.CommonResponse
	AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponseModel
}

// AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponseModel is 空间单元列表带业务属性实例 成功返回结果
type AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitlistwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
