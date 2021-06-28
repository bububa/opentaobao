package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取编辑商品的schema模板 APIResponse
alitrip.item.update.schema.get

获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *AlitripItemUpdateSchemaGetResponse `json:"alitrip_item_update_schema_get_response,omitempty"` 
    AlitripItemUpdateSchemaGetResponse
}

/* model for simplify = false
type AlitripItemUpdateSchemaGetResponse struct {

    // schema数据
    
    SchemaXmlFields   string `json:"schema_xml_fields,omitempty"`
    

}
*/

type AlitripItemUpdateSchemaGetResponse struct {

    // schema数据
    SchemaXmlFields   string `json:"schema_xml_fields,omitempty"`

}
