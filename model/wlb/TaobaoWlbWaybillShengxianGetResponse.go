package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家获取生鲜电子面单号 APIResponse
taobao.wlb.waybill.shengxian.get

商家通过交易订单号获取电子面单接口
*/
type TaobaoWlbWaybillShengxianGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_waybill_shengxian_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 生成是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"