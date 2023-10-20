package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradetemplatequeryAPIResponse 订单服务详情模版查询 API返回值
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
type AlitriptraveltradetemplatequeryAPIResponse struct {
	model.CommonResponse
	AlitriptraveltradetemplatequeryAPIResponseModel
}

// AlitriptraveltradetemplatequeryAPIResponseModel is 订单服务详情模版查询 成功返回结果
type AlitriptraveltradetemplatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单服务标注模版获取结果
	Result *AlitriptraveltradetemplatequeryResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
