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
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_gradeprice_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"