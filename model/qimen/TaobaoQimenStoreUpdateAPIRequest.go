package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreUpdateAPIRequest
门店更新接口 API请求
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息 */
type TaobaoQimenStoreUpdateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 备注
	_remark string
	// 门店主营类目
	_mainCategory int64
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
	// 门店地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
	// 门店类型
	_storeType string
	// 线上门店id
	_storeId int64
	// ERP系统中 门店编码
	_storeCode string
}

// NewTaobaoQimenStoreUpdateRequest 初始化TaobaoQimenStoreUpdateAPIRequest对象
func NewTaobaoQimenStoreUpdateRequest() *TaobaoQimenStoreUpdateAPIRequest {
	return &TaobaoQimenStoreUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreName Setter
// 门店名称
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// Get StoreName Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreName() string {
	return r._storeName
}

// Set is Remark Setter
// 备注
func (r *TaobaoQimenStoreUpdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetRemark() string {
	return r._remark
}

// Set is MainCategory Setter
// 门店主营类目
func (r *TaobaoQimenStoreUpdateAPIRequest) SetMainCategory(_mainCategory int64) error {
	r._mainCategory = _mainCategory
	r.Set("main_category", _mainCategory)
	return nil
}

// Get MainCategory Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetMainCategory() int64 {
	return r._mainCategory
}

// Set is EndTime Setter
// 停止营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is CompanyName Setter
// 商户名称
func (r *TaobaoQimenStoreUpdateAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// Get CompanyName Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetCompanyName() string {
	return r._companyName
}

// Set is StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is StoreStatus Setter
// 门店状态
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreStatus(_storeStatus string) error {
	r._storeStatus = _storeStatus
	r.Set("store_status", _storeStatus)
	return nil
}

// Get StoreStatus Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreStatus() string {
	return r._storeStatus
}

// Set is StoreDescription Setter
// 商户介绍
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreDescription(_storeDescription string) error {
	r._storeDescription = _storeDescription
	r.Set("store_description", _storeDescription)
	return nil
}

// Get StoreDescription Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreDescription() string {
	return r._storeDescription
}

// Set is Address Setter
// 门店地址信息
func (r *TaobaoQimenStoreUpdateAPIRequest) SetAddress(_address *Address) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetAddress() *Address {
	return r._address
}

// Set is ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoQimenStoreUpdateAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetShopId() int64 {
	return r._shopId
}

// Set is StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
	r._storeKeeper = _storeKeeper
	r.Set("store_keeper", _storeKeeper)
	return nil
}

// Get StoreKeeper Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreKeeper() *StoreKeeper {
	return r._storeKeeper
}

// Set is StoreType Setter
// 门店类型
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// Get StoreType Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreType() string {
	return r._storeType
}

// Set is StoreId Setter
// 线上门店id
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is StoreCode Setter
// ERP系统中 门店编码
func (r *TaobaoQimenStoreUpdateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoQimenStoreUpdateAPIRequest) GetStoreCode() string {
	return r._storeCode
}
