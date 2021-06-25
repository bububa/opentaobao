package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务属性查询 APIResponse
taobao.ts.property.get

淘宝服务属性查询
*/
type TaobaoTsPropertyGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTsPropertyGetResponse `json:"taobao_ts_property_get_response,omitempty"`
}

type TaobaoTsPropertyGetResponse struct {

    // 服务收费项相关属性对象
    ServiceItemProperty   *ServiceItemProperty `json:"service_item_property,omitempty"`

}
