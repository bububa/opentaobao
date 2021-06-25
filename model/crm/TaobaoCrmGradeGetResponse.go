package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询等级规则 APIResponse
taobao.crm.grade.get

卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
*/
type TaobaoCrmGradeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGradeGetResponse `json:"taobao_crm_grade_get_response,omitempty"`
}

type TaobaoCrmGradeGetResponse struct {

    // 等级信息集合
    GradePromotions   []GradePromotion `json:"grade_promotions,omitempty"`

}
