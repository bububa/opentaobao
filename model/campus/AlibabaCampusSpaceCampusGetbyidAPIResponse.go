package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacecampusgetbyidAPIResponse 根据园区id获取园区信息 API返回值
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
type AlibabacampusspacecampusgetbyidAPIResponse struct {
	model.CommonResponse
	AlibabacampusspacecampusgetbyidAPIResponseModel
}

// AlibabacampusspacecampusgetbyidAPIResponseModel is 根据园区id获取园区信息 成功返回结果
type AlibabacampusspacecampusgetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_campus_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
