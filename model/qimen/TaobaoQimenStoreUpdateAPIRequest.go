package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoreupdateAPIRequest 门店更新接口 API请求
// taobao.qimen.store.update
//
// 商家在ERP等系统中调用该接口，更新门店信息
type TaobaoqimenstoreupdateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 备注
	_remark string
	// 停止营业时间(只填时，分；只支持半点和整点)
	_endTime string
	// 商户名称
	_companyName string
	// 开始营业时间(只填时，分；只支持半点和整点)
	_startTime string
	// 门店状态
	_storeStatus string
	// 商户介绍
	_storeDescription string
	// 门店类型
	_storeType string
	// ERP系统中 门店编码
	_storeCode string
	// 门店主营类目
	_mainCategory int64
	// 门店地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
	// 线上门店id
	_storeId int64
}

// NewTaobaoqimenstoreupdateRequest 初始化TaobaoqimenstoreupdateAPIRequest对象
func NewTaobaoqimenstoreupdateRequest() *TaobaoqimenstoreupdateAPIRequest {
	return &TaobaoqimenstoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstoreupdateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreName is StoreName Setter
// 门店名称
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// GetStoreName StoreName Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreName() string {
	return r._storeName
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoqimenstoreupdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetRemark() string {
	return r._remark
}

// SetEndTime is EndTime Setter
// 停止营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoqimenstoreupdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCompanyName is CompanyName Setter
// 商户名称
func (r *TaobaoqimenstoreupdateAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetStartTime is StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoqimenstoreupdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetStoreStatus is StoreStatus Setter
// 门店状态
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreStatus(_storeStatus string) error {
	r._storeStatus = _storeStatus
	r.Set("store_status", _storeStatus)
	return nil
}

// GetStoreStatus StoreStatus Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreStatus() string {
	return r._storeStatus
}

// SetStoreDescription is StoreDescription Setter
// 商户介绍
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreDescription(_storeDescription string) error {
	r._storeDescription = _storeDescription
	r.Set("store_description", _storeDescription)
	return nil
}

// GetStoreDescription StoreDescription Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreDescription() string {
	return r._storeDescription
}

// SetStoreType is StoreType Setter
// 门店类型
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// GetStoreType StoreType Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreType() string {
	return r._storeType
}

// SetStoreCode is StoreCode Setter
// ERP系统中 门店编码
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetMainCategory is MainCategory Setter
// 门店主营类目
func (r *TaobaoqimenstoreupdateAPIRequest) SetMainCategory(_mainCategory int64) error {
	r._mainCategory = _mainCategory
	r.Set("main_category", _mainCategory)
	return nil
}

// GetMainCategory MainCategory Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetMainCategory() int64 {
	return r._mainCategory
}

// SetAddress is Address Setter
// 门店地址信息
func (r *TaobaoqimenstoreupdateAPIRequest) SetAddress(_address *Address) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetAddress() *Address {
	return r._address
}

// SetShopId is ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoqimenstoreupdateAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetShopId() int64 {
	return r._shopId
}

// SetStoreKeeper is StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
	r._storeKeeper = _storeKeeper
	r.Set("store_keeper", _storeKeeper)
	return nil
}

// GetStoreKeeper StoreKeeper Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreKeeper() *StoreKeeper {
	return r._storeKeeper
}

// SetStoreId is StoreId Setter
// 线上门店id
func (r *TaobaoqimenstoreupdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoqimenstoreupdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
