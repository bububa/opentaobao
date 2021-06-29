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
    _type   string
    // 发货地区域id
    _fromId   int64
    // 收件人地址
    _toAddress   *AddressDTO
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
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetType() string {
    return r._type
}
// FromId Setter
// 发货地区域id
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetFromId(_fromId int64) error {
    r._fromId = _fromId
    r.Set("from_id", _fromId)
    return nil
}

// FromId Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetFromId() int64 {
    return r._fromId
}
// ToAddress Setter
// 收件人地址
func (r *TaobaoWlbImportThreeplResourceGetRequest) SetToAddress(_toAddress *AddressDTO) error {
    r._toAddress = _toAddress
    r.Set("to_address", _toAddress)
    return nil
}

// ToAddress Getter
func (r TaobaoWlbImportThreeplResourceGetRequest) GetToAddress() *AddressDTO {
    return r._toAddress
}
