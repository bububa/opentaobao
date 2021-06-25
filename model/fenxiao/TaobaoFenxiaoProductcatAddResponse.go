package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增产品线 APIResponse
taobao.fenxiao.productcat.add

新增产品线
*/
type TaobaoFenxiaoProductcatAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductcatAddResponse `json:"taobao_fenxiao_productcat_add_response,omitempty"`
}

type TaobaoFenxiaoProductcatAddResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 产品线ID
    ProductLineId   int64 `json:"product_line_id,omitempty"`

}
