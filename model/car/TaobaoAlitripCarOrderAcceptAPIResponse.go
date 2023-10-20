package car

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderacceptAPIResponse 确认接单 API返回值
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
type TaobaoalitripcarorderacceptAPIResponse struct {
	model.CommonResponse
	TaobaoalitripcarorderacceptAPIResponseModel
}

// TaobaoalitripcarorderacceptAPIResponseModel is 确认接单 成功返回结果
type TaobaoalitripcarorderacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *TaobaoalitripcarorderacceptApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
