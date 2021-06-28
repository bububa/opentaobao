package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订单查询 APIResponse
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录
*/
type CainiaoWaybillPrivacySellerOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_privacy_seller_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误列表
    
    ErrorCodeList   []string `json:"error_code_list,omitempty" xml:"