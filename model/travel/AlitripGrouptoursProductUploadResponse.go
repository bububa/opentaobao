package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
跟团游 产品维护接口 APIResponse
alitrip.grouptours.product.upload

跟团游 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripGrouptoursProductUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlitripGrouptoursProductUploadResponse `json:"alitrip_grouptours_product_upload_response,omitempty"` 
    AlitripGrouptoursProductUploadResponse
}

/* model for simplify = false
type AlitripGrouptoursProductUploadResponse struct {

    // firstResult
    
    FirstResult  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripGrouptoursProductUploadResponse struct {

    // firstResult
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
