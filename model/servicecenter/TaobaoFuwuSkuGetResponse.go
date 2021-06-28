package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取内购服务及SKU详情 APIResponse
taobao.fuwu.sku.get

通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
*/
type TaobaoFuwuSkuGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFuwuSkuGetResponse `json:"fuwu_sku_get_response,omitempty"` 
    TaobaoFuwuSkuGetResponse
}

/* model for simplify = false
type TaobaoFuwuSkuGetResponse struct {

    // 内购服务及SKU详情
    
    Result  *struct {
        ArticleViewResult  *ArticleViewResult `json:"article_view_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFuwuSkuGetResponse struct {

    // 内购服务及SKU详情
    Result   *ArticleViewResult `json:"result,omitempty"`

}
