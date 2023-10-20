package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwlbordergetAPIResponse 根据物流宝订单编号查询物流宝订单概要信息 API返回值
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
type TaobaowlbwlbordergetAPIResponse struct {
	model.CommonResponse
	TaobaowlbwlbordergetAPIResponseModel
}

// TaobaowlbwlbordergetAPIResponseModel is 根据物流宝订单编号查询物流宝订单概要信息 成功返回结果
type TaobaowlbwlbordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wlborder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单对象
	WlbOrder *WlbOrder `json:"wlb_order,omitempty" xml:"wlb_order,omitempty"`
}
