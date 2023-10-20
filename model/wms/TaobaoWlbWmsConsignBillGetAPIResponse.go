package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsConsignBillGetAPIResponse 获取销售订单发货信息 API返回值
// taobao.wlb.wms.consign.bill.get
//
// 获取销售订单发货信息
type TaobaoWlbWmsConsignBillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsConsignBillGetAPIResponseModel
}

// TaobaoWlbWmsConsignBillGetAPIResponseModel is 获取销售订单发货信息 成功返回结果
type TaobaoWlbWmsConsignBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_consign_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品信息列表
	ConsignSendInfoList []Consignsendinfolist `json:"consign_send_info_list,omitempty" xml:"consign_send_info_list>consignsendinfolist,omitempty"`
}
