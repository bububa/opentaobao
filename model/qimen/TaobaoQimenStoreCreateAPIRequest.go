package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstorecreateAPIRequest 门店新增接口 API请求
// taobao.qimen.store.create
//
// isv调用接口来讲线下门店同步到线上
type TaobaoqimenstorecreateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 商户名称
	_companyName string
	// 关闭营业时间(只填时，分；只支持半点和整点)
	_endTime string
	// 开始营业时间(只填时，分；只支持半点和整点)
	_startTime string
	// 门店状态
	_storeStatus string
	// 商户介绍
	_storeDescription string
	// 门店的类型
	_storeType string
	// ERP系统中门店的编码
	_storeCode string
	// 备注
	_remark string
	// 门店主营类目
	_mainCategory int64
	// 地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
}

// NewTaobaoqimenstorecreateRequest 初始化TaobaoqimenstorecreateAPIRequest对象
func NewTaobaoqimenstorecreateRequest() *TaobaoqimenstorecreateAPIRequest {
	return &TaobaoqimenstorecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstorecreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstorecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstorecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreName is StoreName Setter
// 门店名称
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetCompanyName is CompanyName Setter
// 商户名称
func (r *TaobaoqimenstorecreateAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r TaobaoqimenstorecreateAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetEndTime is EndTime Setter
// 关闭营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoqimenstorecreateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoqimenstorecreateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoqimenstorecreateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetStoreStatus is StoreStatus Setter
// 门店状态
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreStatus(_storeStatus string) error {
	r._storeStatus = _storeStatus
	r.Set("store_status", _storeStatus)
	return nil
}

// GetStoreStatus StoreStatus Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreStatus() string {
	return r._storeStatus
}

// SetStoreDescription is StoreDescription Setter
// 商户介绍
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreDescription(_storeDescription string) error {
	r._storeDescription = _storeDescription
	r.Set("store_description", _storeDescription)
	return nil
}

// GetStoreDescription StoreDescription Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreDescription() string {
	return r._storeDescription
}

// SetStoreType is StoreType Setter
// 门店的类型
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreType() string {
	return r._storeType
}

// SetStoreCode is StoreCode Setter
// ERP系统中门店的编码
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoqimenstorecreateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimenstorecreateAPIRequest) GetRemark() string {
	return r._remark
}

// SetMainCategory is MainCategory Setter
// 门店主营类目
func (r *TaobaoqimenstorecreateAPIRequest) SetMainCategory(_mainCategory int64) error {
	r._mainCategory = _mainCategory
	r.Set("main_category", _mainCategory)
	return nil
}

// GetMainCategory MainCategory Getter
func (r TaobaoqimenstorecreateAPIRequest) GetMainCategory() int64 {
	return r._mainCategory
}

// SetAddress is Address Setter
// 地址信息
func (r *TaobaoqimenstorecreateAPIRequest) SetAddress(_address *Address) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoqimenstorecreateAPIRequest) GetAddress() *Address {
	return r._address
}

// SetShopId is ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoqimenstorecreateAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoqimenstorecreateAPIRequest) GetShopId() int64 {
	return r._shopId
}

// SetStoreKeeper is StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoqimenstorecreateAPIRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
	r._storeKeeper = _storeKeeper
	r.Set("store_keeper", _storeKeeper)
	return nil
}

// GetStoreKeeper StoreKeeper Getter
func (r TaobaoqimenstorecreateAPIRequest) GetStoreKeeper() *StoreKeeper {
	return r._storeKeeper
}
