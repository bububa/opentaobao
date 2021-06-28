package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际-清关材料查询 APIResponse
tmall.hk.clearance.get

提供订单收货人身份信息查询功能。
*/
type TmallHkClearanceGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallHkClearanceGetResponse `json:"tmall_hk_clearance_get_response,omitempty"` 
    TmallHkClearanceGetResponse
}

/* model for simplify = false
type TmallHkClearanceGetResponse struct {

    // 查询结果对象
    
    Result  *struct {
        CertifyQueryResult  *CertifyQueryResult `json:"certify_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallHkClearanceGetResponse struct {

    // 查询结果对象
    Result   *CertifyQueryResult `json:"result,omitempty"`

}
