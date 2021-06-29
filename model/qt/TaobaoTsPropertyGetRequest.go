package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务属性查询 API请求
taobao.ts.property.get

淘宝服务属性查询
*/
type TaobaoTsPropertyGetRequest struct {
    model.Params
    // 服务收费项code
    serviceItemCode   string
}

// 初始化TaobaoTsPropertyGetRequest对象
func NewTaobaoTsPropertyGetRequest() *TaobaoTsPropertyGetRequest{
    return &TaobaoTsPropertyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTsPropertyGetRequest) GetApiMethodName() string {
    return "taobao.ts.property.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTsPropertyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceItemCode Setter
// 服务收费项code
func (r *TaobaoTsPropertyGetRequest) SetServiceItemCode(serviceItemCode string) error {
    r.serviceItemCode = serviceItemCode
    r.Set("service_item_code", serviceItemCode)
    return nil
}

// ServiceItemCode Getter
func (r TaobaoTsPropertyGetRequest) GetServiceItemCode() string {
    return r.serviceItemCode
}
