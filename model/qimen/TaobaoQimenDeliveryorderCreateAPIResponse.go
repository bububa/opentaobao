package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryordercreateAPIResponse 发货单创建接口 API返回值
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
type TaobaoqimendeliveryordercreateAPIResponse struct {
	model.CommonResponse
	TaobaoqimendeliveryordercreateAPIResponseModel
}

// TaobaoqimendeliveryordercreateAPIResponseModel is 发货单创建接口 成功返回结果
type TaobaoqimendeliveryordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应字段
	Response *DeliveryOrderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
