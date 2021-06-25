package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家获取生鲜电子面单号 APIResponse
taobao.wlb.waybill.shengxian.get

商家通过交易订单号获取电子面单接口
*/
type TaobaoWlbWaybillShengxianGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillShengxianGetResponse `json:"taobao_wlb_waybill_shengxian_get_response,omitempty"`
}

type TaobaoWlbWaybillShengxianGetResponse struct {

    // 生成是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 成功后返回的生鲜电子面单信息
    FreshWaybill   *FreshWaybill `json:"fresh_waybill,omitempty"`

}
