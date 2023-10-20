package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotspropertygetAPIResponse 淘宝服务属性查询 API返回值
// taobao.ts.property.get
//
// 淘宝服务属性查询
type TaobaotspropertygetAPIResponse struct {
	model.CommonResponse
	TaobaotspropertygetAPIResponseModel
}

// TaobaotspropertygetAPIResponseModel is 淘宝服务属性查询 成功返回结果
type TaobaotspropertygetAPIResponseModel struct {
	XMLName xml.Name `xml:"ts_property_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务收费项相关属性对象
	ServiceItemProperty *ServiceItemProperty `json:"service_item_property,omitempty" xml:"service_item_property,omitempty"`
}
