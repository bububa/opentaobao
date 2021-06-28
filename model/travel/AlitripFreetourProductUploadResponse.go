package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自由行商品发布及编辑接口 APIResponse
alitrip.freetour.product.upload

自由行 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripFreetourProductUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_freetour_product_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // firstResult
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"