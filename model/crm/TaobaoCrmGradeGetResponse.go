package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询等级规则 APIResponse
taobao.crm.grade.get

卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
*/
type TaobaoCrmGradeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_grade_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 等级信息集合
    
    GradePromotions   []GradePromotion `json:"grade_promotions,omitempty" xml:"