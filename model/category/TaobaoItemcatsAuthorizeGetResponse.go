package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商家被授权品牌列表和类目列表 APIResponse
taobao.itemcats.authorize.get

查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
*/
type TaobaoItemcatsAuthorizeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemcatsAuthorizeGetResponse `json:"itemcats_authorize_get_response,omitempty"` 
    TaobaoItemcatsAuthorizeGetResponse
}

/* model for simplify = false
type TaobaoItemcatsAuthorizeGetResponse struct {

    // 里面有3个数组：<br/>Brand[]品牌列表,<br/>ItemCat[] 类目列表<br/>XinpinItemCat[] 针对于C卖家新品类目列表
    
    SellerAuthorize  *struct {
        SellerAuthorize  *SellerAuthorize `json:"seller_authorize,omitempty"`
    } `json:"seller_authorize,omitempty"`
    

}
*/

type TaobaoItemcatsAuthorizeGetResponse struct {

    // 里面有3个数组：<br/>Brand[]品牌列表,<br/>ItemCat[] 类目列表<br/>XinpinItemCat[] 针对于C卖家新品类目列表
    SellerAuthorize   *SellerAuthorize `json:"seller_authorize,omitempty"`

}
