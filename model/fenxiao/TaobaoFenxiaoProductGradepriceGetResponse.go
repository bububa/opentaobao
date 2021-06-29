package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
等级折扣查询 APIResponse
taobao.fenxiao.product.gradeprice.get

等级折扣查询
*/
type TaobaoFenxiaoProductGradepriceGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductGradepriceGetResponse
}

type TaobaoFenxiaoProductGradepriceGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_gradeprice_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 等级折扣列表
    
    GradeDiscounts   []GradeDiscount `json:"grade_discounts,omitempty" xml:"grade_discounts>grade_discount,omitempty"`
    
    
}
