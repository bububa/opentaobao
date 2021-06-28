package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销售订单发货信息 APIResponse
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息
*/
type TaobaoWlbWmsConsignBillGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_consign_bill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品信息列表
    
    ConsignSendInfoList   []Consignsendinfolist `json:"consign_send_info_list,omitempty" xml:"