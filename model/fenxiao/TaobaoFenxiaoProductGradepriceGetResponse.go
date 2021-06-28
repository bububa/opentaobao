package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
等级折扣查询 APIResponse
taobao.fenxiao.product.gradeprice.get

等级折扣查询
*/
type TaobaoFenxiaoProductGradepriceGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductGradepriceGetResponse `json:"fenxiao_product_gradeprice_get_response,omitempty"` 
    TaobaoFenxiaoProductGradepriceGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductGradepriceGetResponse struct {

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 等级折扣列表
    
    GradeDiscounts  struct {
        GradeDiscount  []GradeDiscount `json:"grade_discount,omitempty"`
    } `json:"grade_discounts,omitempty"`
    

}
*/

type TaobaoFenxiaoProductGradepriceGetResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 等级折扣列表
    GradeDiscounts   []GradeDiscount `json:"grade_discounts,omitempty"`

}
