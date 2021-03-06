package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceTypeGetpageresultAPIResponse 分页查询空间类别接口 API返回值
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
type AlibabaCampusSpaceTypeGetpageresultAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceTypeGetpageresultAPIResponseModel
}

// AlibabaCampusSpaceTypeGetpageresultAPIResponseModel is 分页查询空间类别接口 成功返回结果
type AlibabaCampusSpaceTypeGetpageresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_type_getpageresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
