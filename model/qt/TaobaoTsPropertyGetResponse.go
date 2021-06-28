package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务属性查询 APIResponse
taobao.ts.property.get

淘宝服务属性查询
*/
type TaobaoTsPropertyGetAPIResponse struct {
    model.CommonResponse
    TaobaoTsPropertyGetResponse
}

type TaobaoTsPropertyGetResponse struct {
    XMLName xml.Name `xml:"ts_property_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务收费项相关属性对象
    
    ServiceItemProperty   *ServiceItemProperty `json:"service_item_property,omitempty" xml:"service_item_property,omitempty"`

    
}
