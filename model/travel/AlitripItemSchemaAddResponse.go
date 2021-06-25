package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
使用schema模板进行商品发布 APIResponse
alitrip.item.schema.add

飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemSchemaAddAPIResponse struct {
    model.CommonResponse
    Response *AlitripItemSchemaAddResponse `json:"alitrip_item_schema_add_response,omitempty"`
}

type AlitripItemSchemaAddResponse struct {

    // 结果
    Result   *TopTravelItem `json:"result,omitempty"`

}
