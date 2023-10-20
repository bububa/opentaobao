package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsconsignordercreateandsendAPIResponse 创建订单并发货 API返回值
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
type TaobaologisticsconsignordercreateandsendAPIResponse struct {
	model.CommonResponse
	TaobaologisticsconsignordercreateandsendAPIResponseModel
}

// TaobaologisticsconsignordercreateandsendAPIResponseModel is 创建订单并发货 成功返回结果
type TaobaologisticsconsignordercreateandsendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_consign_order_createandsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
