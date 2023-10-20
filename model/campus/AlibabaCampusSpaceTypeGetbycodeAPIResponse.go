package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacetypegetbycodeAPIResponse 根据类别编码查询类别 API返回值
// alibaba.campus.space.type.getbycode
//
// 根据类别编码查询类别
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getByCode
type AlibabacampusspacetypegetbycodeAPIResponse struct {
	model.CommonResponse
	AlibabacampusspacetypegetbycodeAPIResponseModel
}

// AlibabacampusspacetypegetbycodeAPIResponseModel is 根据类别编码查询类别 成功返回结果
type AlibabacampusspacetypegetbycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_type_getbycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
