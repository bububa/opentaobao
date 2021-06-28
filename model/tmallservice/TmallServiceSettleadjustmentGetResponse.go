package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询结算调整单单条记录 APIResponse
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息
*/
type TmallServiceSettleadjustmentGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallServiceSettleadjustmentGetResponse `json:"tmall_service_settleadjustment_get_response,omitempty"` 
    TmallServiceSettleadjustmentGetResponse
}

/* model for simplify = false
type TmallServiceSettleadjustmentGetResponse struct {

    // result
    
    Result  *struct {
        TmallServiceSettleadjustmentGetResult  *TmallServiceSettleadjustmentGetResult `json:"tmall_service_settleadjustment_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServiceSettleadjustmentGetResponse struct {

    // result
    Result   *TmallServiceSettleadjustmentGetResult `json:"result,omitempty"`

}
