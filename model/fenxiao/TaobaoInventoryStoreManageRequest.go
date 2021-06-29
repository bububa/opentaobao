package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建或更新仓库 API请求
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryStoreManageRequest struct {
    model.Params
    // 参数定义，ADD：新建; UPDATE：更新
    _operateType   string
    // 商家的仓库编码，不允许重复，不允许更新
    _storeCode   string
    // 商家的仓库名称，可更新
    _storeName   string
    // 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
    _storeType   string
    // 仓库简称，可更新
    _aliasName   string
    // 仓库的物理地址，可更新
    _address   string
    // 仓库区域名，可更新
    _addressAreaName   string
    // 联系人，可更新
    _contact   string
    // 联系电话，可更新
    _phone   string
    // 邮编，可更新
    _postcode   int64
}

// 初始化TaobaoInventoryStoreManageRequest对象
func NewTaobaoInventoryStoreManageRequest() *TaobaoInventoryStoreManageRequest{
    return &TaobaoInventoryStoreManageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryStoreManageRequest) GetApiMethodName() string {
    return "taobao.inventory.store.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryStoreManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperateType Setter
// 参数定义，ADD：新建; UPDATE：更新
func (r *TaobaoInventoryStoreManageRequest) SetOperateType(_operateType string) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoInventoryStoreManageRequest) GetOperateType() string {
    return r._operateType
}
// StoreCode Setter
// 商家的仓库编码，不允许重复，不允许更新
func (r *TaobaoInventoryStoreManageRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreCode() string {
    return r._storeCode
}
// StoreName Setter
// 商家的仓库名称，可更新
func (r *TaobaoInventoryStoreManageRequest) SetStoreName(_storeName string) error {
    r._storeName = _storeName
    r.Set("store_name", _storeName)
    return nil
}

// StoreName Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreName() string {
    return r._storeName
}
// StoreType Setter
// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
func (r *TaobaoInventoryStoreManageRequest) SetStoreType(_storeType string) error {
    r._storeType = _storeType
    r.Set("store_type", _storeType)
    return nil
}

// StoreType Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreType() string {
    return r._storeType
}
// AliasName Setter
// 仓库简称，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAliasName(_aliasName string) error {
    r._aliasName = _aliasName
    r.Set("alias_name", _aliasName)
    return nil
}

// AliasName Getter
func (r TaobaoInventoryStoreManageRequest) GetAliasName() string {
    return r._aliasName
}
// Address Setter
// 仓库的物理地址，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoInventoryStoreManageRequest) GetAddress() string {
    return r._address
}
// AddressAreaName Setter
// 仓库区域名，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAddressAreaName(_addressAreaName string) error {
    r._addressAreaName = _addressAreaName
    r.Set("address_area_name", _addressAreaName)
    return nil
}

// AddressAreaName Getter
func (r TaobaoInventoryStoreManageRequest) GetAddressAreaName() string {
    return r._addressAreaName
}
// Contact Setter
// 联系人，可更新
func (r *TaobaoInventoryStoreManageRequest) SetContact(_contact string) error {
    r._contact = _contact
    r.Set("contact", _contact)
    return nil
}

// Contact Getter
func (r TaobaoInventoryStoreManageRequest) GetContact() string {
    return r._contact
}
// Phone Setter
// 联系电话，可更新
func (r *TaobaoInventoryStoreManageRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TaobaoInventoryStoreManageRequest) GetPhone() string {
    return r._phone
}
// Postcode Setter
// 邮编，可更新
func (r *TaobaoInventoryStoreManageRequest) SetPostcode(_postcode int64) error {
    r._postcode = _postcode
    r.Set("postcode", _postcode)
    return nil
}

// Postcode Getter
func (r TaobaoInventoryStoreManageRequest) GetPostcode() int64 {
    return r._postcode
}
