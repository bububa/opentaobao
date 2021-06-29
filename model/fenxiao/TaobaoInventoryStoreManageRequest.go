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
    operateType   string
    // 商家的仓库编码，不允许重复，不允许更新
    storeCode   string
    // 商家的仓库名称，可更新
    storeName   string
    // 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
    storeType   string
    // 仓库简称，可更新
    aliasName   string
    // 仓库的物理地址，可更新
    address   string
    // 仓库区域名，可更新
    addressAreaName   string
    // 联系人，可更新
    contact   string
    // 联系电话，可更新
    phone   string
    // 邮编，可更新
    postcode   int64
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
func (r *TaobaoInventoryStoreManageRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

// OperateType Getter
func (r TaobaoInventoryStoreManageRequest) GetOperateType() string {
    return r.operateType
}
// StoreCode Setter
// 商家的仓库编码，不允许重复，不允许更新
func (r *TaobaoInventoryStoreManageRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreCode() string {
    return r.storeCode
}
// StoreName Setter
// 商家的仓库名称，可更新
func (r *TaobaoInventoryStoreManageRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

// StoreName Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreName() string {
    return r.storeName
}
// StoreType Setter
// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
func (r *TaobaoInventoryStoreManageRequest) SetStoreType(storeType string) error {
    r.storeType = storeType
    r.Set("store_type", storeType)
    return nil
}

// StoreType Getter
func (r TaobaoInventoryStoreManageRequest) GetStoreType() string {
    return r.storeType
}
// AliasName Setter
// 仓库简称，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAliasName(aliasName string) error {
    r.aliasName = aliasName
    r.Set("alias_name", aliasName)
    return nil
}

// AliasName Getter
func (r TaobaoInventoryStoreManageRequest) GetAliasName() string {
    return r.aliasName
}
// Address Setter
// 仓库的物理地址，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoInventoryStoreManageRequest) GetAddress() string {
    return r.address
}
// AddressAreaName Setter
// 仓库区域名，可更新
func (r *TaobaoInventoryStoreManageRequest) SetAddressAreaName(addressAreaName string) error {
    r.addressAreaName = addressAreaName
    r.Set("address_area_name", addressAreaName)
    return nil
}

// AddressAreaName Getter
func (r TaobaoInventoryStoreManageRequest) GetAddressAreaName() string {
    return r.addressAreaName
}
// Contact Setter
// 联系人，可更新
func (r *TaobaoInventoryStoreManageRequest) SetContact(contact string) error {
    r.contact = contact
    r.Set("contact", contact)
    return nil
}

// Contact Getter
func (r TaobaoInventoryStoreManageRequest) GetContact() string {
    return r.contact
}
// Phone Setter
// 联系电话，可更新
func (r *TaobaoInventoryStoreManageRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r TaobaoInventoryStoreManageRequest) GetPhone() string {
    return r.phone
}
// Postcode Setter
// 邮编，可更新
func (r *TaobaoInventoryStoreManageRequest) SetPostcode(postcode int64) error {
    r.postcode = postcode
    r.Set("postcode", postcode)
    return nil
}

// Postcode Getter
func (r TaobaoInventoryStoreManageRequest) GetPostcode() int64 {
    return r.postcode
}
