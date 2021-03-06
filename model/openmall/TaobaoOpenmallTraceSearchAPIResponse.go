package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTraceSearchAPIResponse 获取Openmall订单物流流转信息 API返回值
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
type TaobaoOpenmallTraceSearchAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTraceSearchAPIResponseModel
}

// TaobaoOpenmallTraceSearchAPIResponseModel is 获取Openmall订单物流流转信息 成功返回结果
type TaobaoOpenmallTraceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trace_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopLogisticsDetailTraceVo `json:"result,omitempty" xml:"result,omitempty"`
}
