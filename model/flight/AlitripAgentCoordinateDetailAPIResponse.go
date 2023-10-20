package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinatedetailAPIResponse 商家协同单查询详情 API返回值
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
type AlitripagentcoordinatedetailAPIResponse struct {
	model.CommonResponse
	AlitripagentcoordinatedetailAPIResponseModel
}

// AlitripagentcoordinatedetailAPIResponseModel is 商家协同单查询详情 成功返回结果
type AlitripagentcoordinatedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询协同单详情返回内容
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
