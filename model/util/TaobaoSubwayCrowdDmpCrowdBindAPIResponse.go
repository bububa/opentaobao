package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdDmpCrowdBindAPIResponse 直通车绑定达摩盘人群 API返回值
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
type TaobaoSubwayCrowdDmpCrowdBindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel
}

// TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel is 直通车绑定达摩盘人群 成功返回结果
type TaobaoSubwayCrowdDmpCrowdBindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_crowd_dmp_crowd_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定结果数据
	Result []CrowdBindTopResultDto `json:"result,omitempty" xml:"result>crowd_bind_top_result_dto,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
