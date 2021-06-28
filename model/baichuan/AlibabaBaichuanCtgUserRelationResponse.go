package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户 APIResponse
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系
*/
type AlibabaBaichuanCtgUserRelationAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBaichuanCtgUserRelationResponse `json:"alibaba_baichuan_ctg_user_relation_response,omitempty"` 
    AlibabaBaichuanCtgUserRelationResponse
}

/* model for simplify = false
type AlibabaBaichuanCtgUserRelationResponse struct {

    // 返回的整体结果
    
    Result  *struct {
        AlibabaBaichuanCtgUserRelationResult  *AlibabaBaichuanCtgUserRelationResult `json:"alibaba_baichuan_ctg_user_relation_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaBaichuanCtgUserRelationResponse struct {

    // 返回的整体结果
    Result   *AlibabaBaichuanCtgUserRelationResult `json:"result,omitempty"`

}
