package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripitemaddschemagetAPIResponse 获取商品发布模板 API返回值
// alitrip.item.add.schema.get
//
// 发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripitemaddschemagetAPIResponse struct {
	model.CommonResponse
	AlitripitemaddschemagetAPIResponseModel
}

// AlitripitemaddschemagetAPIResponseModel is 获取商品发布模板 成功返回结果
type AlitripitemaddschemagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_item_add_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// schema模板数据
	SchemaXmlFields string `json:"schema_xml_fields,omitempty" xml:"schema_xml_fields,omitempty"`
}
