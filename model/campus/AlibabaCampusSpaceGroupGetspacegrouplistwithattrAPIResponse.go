package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponse 分页查询空间分组业务属性 API返回值
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
type AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponse struct {
	model.CommonResponse
	AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponseModel
}

// AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponseModel is 分页查询空间分组业务属性 成功返回结果
type AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegrouplistwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
