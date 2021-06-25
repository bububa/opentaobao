package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
使用schema进行商品编辑 APIResponse
alitrip.item.schema.update

飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemSchemaUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlitripItemSchemaUpdateResponse `json:"alitrip_item_schema_update_response,omitempty"`
}

type AlitripItemSchemaUpdateResponse struct {

    // 结果
    Result   *TopTravelItem `json:"result,omitempty"`

}
