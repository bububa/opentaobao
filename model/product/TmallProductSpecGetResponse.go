package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据产品规格的Id号获取当个的规格信息 APIResponse
tmall.product.spec.get

通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核<br/>通过参看这个ProductSpec的status判断：<br/>1:表示审核通过<br/>3:表示等待审核。<br/>如果你的id找不到数据，那么就是审核被拒绝。
*/
type TmallProductSpecGetAPIResponse struct {
    model.CommonResponse
    Response *TmallProductSpecGetResponse `json:"tmall_product_spec_get_response,omitempty"`
}

type TmallProductSpecGetResponse struct {

    // 返回的产品规格信息，注意，这个产品规格信息可能是等待审核的，不一定可用。根据状态判断1：表示审核通过<br/>3：表示等待审核。
    ProductSpec   *ProductSpec `json:"product_spec,omitempty"`

}
