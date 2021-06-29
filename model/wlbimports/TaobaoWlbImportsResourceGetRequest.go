package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有服务列表 API请求
taobao.wlb.imports.resource.get

一般进口TOP接口，获取所有服务列表。
*/
type TaobaoWlbImportsResourceGetRequest struct {
    model.Params
    // 卖家发货地区域ID
    fromId   int64
    // 买家收货地信息
    toAddress   *ReciverAddressDo
}

// 初始化TaobaoWlbImportsResourceGetRequest对象
func NewTaobaoWlbImportsResourceGetRequest() *TaobaoWlbImportsResourceGetRequest{
    return &TaobaoWlbImportsResourceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsResourceGetRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.resource.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsResourceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FromId Setter
// 卖家发货地区域ID
func (r *TaobaoWlbImportsResourceGetRequest) SetFromId(fromId int64) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

// FromId Getter
func (r TaobaoWlbImportsResourceGetRequest) GetFromId() int64 {
    return r.fromId
}
// ToAddress Setter
// 买家收货地信息
func (r *TaobaoWlbImportsResourceGetRequest) SetToAddress(toAddress *ReciverAddressDo) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

// ToAddress Getter
func (r TaobaoWlbImportsResourceGetRequest) GetToAddress() *ReciverAddressDo {
    return r.toAddress
}
