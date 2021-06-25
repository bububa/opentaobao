package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建分销和后端商品映射关系 APIResponse
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。
*/
type TaobaoFenxiaoProductMapAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductMapAddResponse `json:"taobao_fenxiao_product_map_add_response,omitempty"`
}

type TaobaoFenxiaoProductMapAddResponse struct {

    // 操作结果
    IsSuccess   bool `json:"is_success,omitempty"`

}
