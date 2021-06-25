package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商15分钟获取数据 APIResponse
tmall.service.settleadjustment.search

天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
*/
type TmallServiceSettleadjustmentSearchAPIResponse struct {
    model.CommonResponse
    Response *TmallServiceSettleadjustmentSearchResponse `json:"tmall_service_settleadjustment_search_response,omitempty"`
}

type TmallServiceSettleadjustmentSearchResponse struct {

    // result
    Result   *TmallServiceSettleadjustmentSearchResult `json:"result,omitempty"`

}
