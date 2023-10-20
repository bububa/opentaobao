package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventorystoremanageAPIRequest 创建或更新仓库 API请求
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
type TaobaoinventorystoremanageAPIRequest struct {
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
	// 联系人，可更新
	_contact string
	// 联系电话，可更新
	_phone string
	// 仓库区域名，可更新
	_addressAreaName string
	// 邮编，可更新
	_postcode int64
}

// NewTaobaoinventorystoremanageRequest 初始化TaobaoinventorystoremanageAPIRequest对象
func NewTaobaoinventorystoremanageRequest() *TaobaoinventorystoremanageAPIRequest {
	return &TaobaoinventorystoremanageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventorystoremanageAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.store.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventorystoremanageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventorystoremanageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperateType is OperateType Setter
// 参数定义，ADD：新建; UPDATE：更新
func (r *TaobaoinventorystoremanageAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoinventorystoremanageAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetStoreCode is StoreCode Setter
// 商家的仓库编码，不允许重复，不允许更新
func (r *TaobaoinventorystoremanageAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoinventorystoremanageAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetStoreName is StoreName Setter
// 商家的仓库名称，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r TaobaoinventorystoremanageAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetStoreType is StoreType Setter
// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
func (r *TaobaoinventorystoremanageAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r TaobaoinventorystoremanageAPIRequest) GetStoreType() string {
	return r._storeType
}

// SetAliasName is AliasName Setter
// 仓库简称，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetAliasName(_aliasName string) error {
	r._aliasName = _aliasName
	r.Set("alias_name", _aliasName)
	return nil
}

// GetAliasName AliasName Getter
func (r TaobaoinventorystoremanageAPIRequest) GetAliasName() string {
	return r._aliasName
}

// SetAddress is Address Setter
// 仓库的物理地址，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoinventorystoremanageAPIRequest) GetAddress() string {
	return r._address
}

// SetContact is Contact Setter
// 联系人，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetContact(_contact string) error {
	r._contact = _contact
	r.Set("contact", _contact)
	return nil
}

// GetContact Contact Getter
func (r TaobaoinventorystoremanageAPIRequest) GetContact() string {
	return r._contact
}

// SetPhone is Phone Setter
// 联系电话，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaoinventorystoremanageAPIRequest) GetPhone() string {
	return r._phone
}

// SetAddressAreaName is AddressAreaName Setter
// 仓库区域名，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetAddressAreaName(_addressAreaName string) error {
	r._addressAreaName = _addressAreaName
	r.Set("address_area_name", _addressAreaName)
	return nil
}

// GetAddressAreaName AddressAreaName Getter
func (r TaobaoinventorystoremanageAPIRequest) GetAddressAreaName() string {
	return r._addressAreaName
}

// SetPostcode is Postcode Setter
// 邮编，可更新
func (r *TaobaoinventorystoremanageAPIRequest) SetPostcode(_postcode int64) error {
	r._postcode = _postcode
	r.Set("postcode", _postcode)
	return nil
}

// GetPostcode Postcode Getter
func (r TaobaoinventorystoremanageAPIRequest) GetPostcode() int64 {
	return r._postcode
}
