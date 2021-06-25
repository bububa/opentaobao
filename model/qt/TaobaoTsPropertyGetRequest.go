package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务属性查询 APIRequest
taobao.ts.property.get

淘宝服务属性查询
*/
type TaobaoTsPropertyGetRequest struct {
    model.Params

    // 服务收费项code
    serviceItemCode   string 

}

func NewTaobaoTsPropertyGetRequest() *TaobaoTsPropertyGetRequest{
    return &TaobaoTsPropertyGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTsPropertyGetRequest) GetApiMethodName() string {
    return "taobao.ts.property.get"
}

func (r TaobaoTsPropertyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTsPropertyGetRequest) SetServiceItemCode(serviceItemCode string) error {
    r.serviceItemCode = serviceItemCode
    r.Set("service_item_code", serviceItemCode)
    return nil
}

func (r TaobaoTsPropertyGetRequest) GetServiceItemCode() string {
    return r.serviceItemCode
}

