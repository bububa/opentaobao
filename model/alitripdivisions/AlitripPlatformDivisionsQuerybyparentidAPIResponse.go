package alitripdivisions

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformdivisionsquerybyparentidAPIResponse 根据父节点id查询下级行政区划数据 API返回值
// alitrip.platform.divisions.querybyparentid
//
// 根据行政区划id查询下一层级行政区划数据
type AlitripplatformdivisionsquerybyparentidAPIResponse struct {
	model.CommonResponse
	AlitripplatformdivisionsquerybyparentidAPIResponseModel
}

// AlitripplatformdivisionsquerybyparentidAPIResponseModel is 根据父节点id查询下级行政区划数据 成功返回结果
type AlitripplatformdivisionsquerybyparentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_divisions_querybyparentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripplatformdivisionsquerybyparentidResult `json:"result,omitempty" xml:"result,omitempty"`
}
