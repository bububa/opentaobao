package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据指定的资源获取所有中转仓列表 API请求
taobao.wlb.imports.resource.transferstore.get

根据指定的资源获取所有中转仓列表
*/
type TaobaoWlbImportsResourceTransferstoreGetRequest struct {
    model.Params
    // 通过taobao.wlb.imports.resource.get接口查询出来的资源ID
    _resourceId   int64
    // 卖家发货地址的区域ID，如果不填则为默认发货地址ID
    _fromId   int64
    // 商品前台叶子类目ID
    _cids   []int64
    // 买家收货地信息
    _toAddress   *ReciverAddressDo
}

// 初始化TaobaoWlbImportsResourceTransferstoreGetRequest对象
func NewTaobaoWlbImportsResourceTransferstoreGetRequest() *TaobaoWlbImportsResourceTransferstoreGetRequest{
    return &TaobaoWlbImportsResourceTransferstoreGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.resource.transferstore.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResourceId Setter
// 通过taobao.wlb.imports.resource.get接口查询出来的资源ID
func (r *TaobaoWlbImportsResourceTransferstoreGetRequest) SetResourceId(_resourceId int64) error {
    r._resourceId = _resourceId
    r.Set("resource_id", _resourceId)
    return nil
}

// ResourceId Getter
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetResourceId() int64 {
    return r._resourceId
}
// FromId Setter
// 卖家发货地址的区域ID，如果不填则为默认发货地址ID
func (r *TaobaoWlbImportsResourceTransferstoreGetRequest) SetFromId(_fromId int64) error {
    r._fromId = _fromId
    r.Set("from_id", _fromId)
    return nil
}

// FromId Getter
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetFromId() int64 {
    return r._fromId
}
// Cids Setter
// 商品前台叶子类目ID
func (r *TaobaoWlbImportsResourceTransferstoreGetRequest) SetCids(_cids []int64) error {
    r._cids = _cids
    r.Set("cids", _cids)
    return nil
}

// Cids Getter
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetCids() []int64 {
    return r._cids
}
// ToAddress Setter
// 买家收货地信息
func (r *TaobaoWlbImportsResourceTransferstoreGetRequest) SetToAddress(_toAddress *ReciverAddressDo) error {
    r._toAddress = _toAddress
    r.Set("to_address", _toAddress)
    return nil
}

// ToAddress Getter
func (r TaobaoWlbImportsResourceTransferstoreGetRequest) GetToAddress() *ReciverAddressDo {
    return r._toAddress
}
