package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取编辑商品的schema模板 APIResponse
alitrip.item.update.schema.get

获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_item_update_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // schema数据
    
    SchemaXmlFields   string `json:"schema_xml_fields,omitempty" xml:"