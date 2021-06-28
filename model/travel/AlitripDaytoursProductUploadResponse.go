package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
境外一日游/多日游 产品维护接口 APIResponse
alitrip.daytours.product.upload

境外一日游/多日游 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripDaytoursProductUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlitripDaytoursProductUploadResponse `json:"alitrip_daytours_product_upload_response,omitempty"` 
    AlitripDaytoursProductUploadResponse
}

/* model for simplify = false
type AlitripDaytoursProductUploadResponse struct {

    // 商品维护结果
    
    FirstResult  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripDaytoursProductUploadResponse struct {

    // 商品维护结果
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
