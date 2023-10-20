package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetspaceunitwithattrAPIResponse 空间单元id查业务属性实例 API返回值
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
type AlibabacampusspaceunitgetspaceunitwithattrAPIResponse struct {
	model.CommonResponse
	AlibabacampusspaceunitgetspaceunitwithattrAPIResponseModel
}

// AlibabacampusspaceunitgetspaceunitwithattrAPIResponseModel is 空间单元id查业务属性实例 成功返回结果
type AlibabacampusspaceunitgetspaceunitwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
