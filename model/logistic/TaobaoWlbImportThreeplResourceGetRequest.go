package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮获取资源列表 API请求
taobao.wlb.import.threepl.resource.get

获取3pl直邮的发货可用资源
*/
type TaobaoWlbImportThreeplResourceGetRequest struct {
    model.Params
    // ONLINE或者OFFLINE
    type   string
    // 发货地区域id
    fromId   int64
    // 收件人地址
    toAddress   *AddressDto
}

// 初始化TaobaoWlbImportThreeplResourceGetRequest对象
func NewTaobaoWlbImportThreeplResourceGetRequest() *TaobaoWlbImportThreeplResourceGetRequest{
    return &TaobaoWlbImportThreeplResourceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportThreeplResourceGetRequest) GetApiMethodName() string {
    return "taobao.wlb.import.threepl.resource.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportThreeplResourceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// ONLINE或者OFFLINE
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetType() string {
    return r.type
}
// FromId Setter
// 发货地区域id
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetFromId(fromId int64) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

// FromId Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetFromId() int64 {
    return r.fromId
}
// ToAddress Setter
// 收件人地址
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetToAddress(toAddress *AddressDto) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

// ToAddress Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetToAddress() *AddressDto {
    return r.toAddress
}
