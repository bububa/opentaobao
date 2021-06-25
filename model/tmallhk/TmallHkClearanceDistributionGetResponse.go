package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销供应商获取清关材料 APIResponse
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
*/
type TmallHkClearanceDistributionGetAPIResponse struct {
    model.CommonResponse
    Response *TmallHkClearanceDistributionGetResponse `json:"tmall_hk_clearance_distribution_get_response,omitempty"`
}

type TmallHkClearanceDistributionGetResponse struct {

    // 查询结果对象
    Result   *CertifyQueryResult `json:"result,omitempty"`

}
