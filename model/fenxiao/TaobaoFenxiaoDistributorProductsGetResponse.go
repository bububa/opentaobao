package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询产品信息 APIResponse
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息
*/
type TaobaoFenxiaoDistributorProductsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoDistributorProductsGetResponse `json:"taobao_fenxiao_distributor_products_get_response,omitempty"`
}

type TaobaoFenxiaoDistributorProductsGetResponse struct {

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    Products   []FenxiaoProduct `json:"products,omitempty"`

}
