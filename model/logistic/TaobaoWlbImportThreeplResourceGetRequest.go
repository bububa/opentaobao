package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮获取资源列表 APIRequest
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

func NewTaobaoWlbImportThreeplResourceGetRequest() *TaobaoWlbImportThreeplResourceGetRequest{
    return &TaobaoWlbImportThreeplResourceGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportThreeplResourceGetRequest) GetApiMethodName() string {
    return "taobao.wlb.import.threepl.resource.get"
}

func (r TaobaoWlbImportThreeplResourceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportThreeplResourceGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbImportThreeplResourceGetRequest) GetType() string {
    return r.type
}

func (r *TaobaoWlbImportThreeplResourceGetRequest) SetFromId(fromId int64) error {
    r.fromId = fromId
    r.Set("from_id", fromId)
    return nil
}

func (r TaobaoWlbImportThreeplResourceGetRequest) GetFromId() int64 {
    return r.fromId
}

func (r *TaobaoWlbImportThreeplResourceGetRequest) SetToAddress(toAddress *AddressDto) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

func (r TaobaoWlbImportThreeplResourceGetRequest) GetToAddress() *AddressDto {
    return r.toAddress
}

