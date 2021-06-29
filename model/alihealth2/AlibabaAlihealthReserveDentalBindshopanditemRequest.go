package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定门店信息，商品信息 API请求
alibaba.alihealth.reserve.dental.bindshopanditem

绑定门店信息，商品信息
*/
type AlibabaAlihealthReserveDentalBindshopanditemRequest struct {
    model.Params
    // bind_list
    _bindList   []BindDTO
}

// 初始化AlibabaAlihealthReserveDentalBindshopanditemRequest对象
func NewAlibabaAlihealthReserveDentalBindshopanditemRequest() *AlibabaAlihealthReserveDentalBindshopanditemRequest{
    return &AlibabaAlihealthReserveDentalBindshopanditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.bindshopanditem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BindList Setter
// bind_list
func (r *AlibabaAlihealthReserveDentalBindshopanditemRequest) SetBindList(_bindList []BindDTO) error {
    r._bindList = _bindList
    r.Set("bind_list", _bindList)
    return nil
}

// BindList Getter
func (r AlibabaAlihealthReserveDentalBindshopanditemRequest) GetBindList() []BindDTO {
    return r._bindList
}
