package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryStoreManageAPIRequest
创建或更新仓库 API请求
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息 */
type TaobaoInventoryStoreManageAPIRequest struct {
	model.Params
	// 参数定义，ADD：新建; UPDATE：更新
	_operateType string
	// 商家的仓库编码，不允许重复，不允许更新
	_storeCode string
	// 商家的仓库名称，可更新
	_storeName string
	// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
	_storeType string
	// 仓库简称，可更新
	_aliasName string
	// 仓库的物理地址，可更新
	_address string
	// 仓库区域名，可更新
	_addressAreaName string
	// 联系人，可更新
	_contact string
	// 联系电话，可更新
	_phone string
	// 邮编，可更新
	_postcode int64
}

// NewTaobaoInventoryStoreManageRequest 初始化TaobaoInventoryStoreManageAPIRequest对象
func NewTaobaoInventoryStoreManageRequest() *TaobaoInventoryStoreManageAPIRequest {
	return &TaobaoInventoryStoreManageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryStoreManageAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.store.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryStoreManageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OperateType Setter
// 参数定义，ADD：新建; UPDATE：更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// Get OperateType Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetOperateType() string {
	return r._operateType
}

// Set is StoreCode Setter
// 商家的仓库编码，不允许重复，不允许更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// Set is StoreName Setter
// 商家的仓库名称，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// Get StoreName Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetStoreName() string {
	return r._storeName
}

// Set is StoreType Setter
// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
func (r *TaobaoInventoryStoreManageAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// Get StoreType Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetStoreType() string {
	return r._storeType
}

// Set is AliasName Setter
// 仓库简称，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetAliasName(_aliasName string) error {
	r._aliasName = _aliasName
	r.Set("alias_name", _aliasName)
	return nil
}

// Get AliasName Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetAliasName() string {
	return r._aliasName
}

// Set is Address Setter
// 仓库的物理地址，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetAddress() string {
	return r._address
}

// Set is AddressAreaName Setter
// 仓库区域名，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetAddressAreaName(_addressAreaName string) error {
	r._addressAreaName = _addressAreaName
	r.Set("address_area_name", _addressAreaName)
	return nil
}

// Get AddressAreaName Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetAddressAreaName() string {
	return r._addressAreaName
}

// Set is Contact Setter
// 联系人，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetContact(_contact string) error {
	r._contact = _contact
	r.Set("contact", _contact)
	return nil
}

// Get Contact Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetContact() string {
	return r._contact
}

// Set is Phone Setter
// 联系电话，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetPhone() string {
	return r._phone
}

// Set is Postcode Setter
// 邮编，可更新
func (r *TaobaoInventoryStoreManageAPIRequest) SetPostcode(_postcode int64) error {
	r._postcode = _postcode
	r.Set("postcode", _postcode)
	return nil
}

// Get Postcode Getter
func (r TaobaoInventoryStoreManageAPIRequest) GetPostcode() int64 {
	return r._postcode
}
