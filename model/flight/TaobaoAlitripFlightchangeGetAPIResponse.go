package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripFlightchangeGetAPIResponse 获取航变信息 API返回值
// taobao.alitrip.flightchange.get
//
// 查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
type TaobaoAlitripFlightchangeGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripFlightchangeGetAPIResponseModel
}

// TaobaoAlitripFlightchangeGetAPIResponseModel is 获取航变信息 成功返回结果
type TaobaoAlitripFlightchangeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flightchange_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	ResultDO *TaobaoAlitripFlightchangeGetResultDo `json:"result_d_o,omitempty" xml:"result_d_o,omitempty"`
}
