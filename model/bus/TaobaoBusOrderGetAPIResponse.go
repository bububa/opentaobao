package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusOrderGetAPIResponse 汽车票订单查询 API返回值
// taobao.bus.order.get
//
// 商家汽车票订单查询
type TaobaoBusOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusOrderGetAPIResponseModel
}

// TaobaoBusOrderGetAPIResponseModel is 汽车票订单查询 成功返回结果
type TaobaoBusOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单查询返回对象
	Result *B2BOrderQueryRp `json:"result,omitempty" xml:"result,omitempty"`
}
