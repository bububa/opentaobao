package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreCreateAPIRequest
门店新增接口 API请求
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上 */
type TaobaoQimenStoreCreateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 门店主营类目
	_mainCategory int64
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
	// 地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
	// 门店的类型
	_storeType string
	// ERP系统中门店的编码
	_storeCode string
	// 备注
	_remark string
}

// NewTaobaoQimenStoreCreateRequest 初始化TaobaoQimenStoreCreateAPIRequest对象
func NewTaobaoQimenStoreCreateRequest() *TaobaoQimenStoreCreateAPIRequest {
	return &TaobaoQimenStoreCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreName Setter
// 门店名称
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreName(_storeName string) error {
	r._storeName = _storeName
	r.Set("store_name", _storeName)
	return nil
}

// Get StoreName Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreName() string {
	return r._storeName
}

// Set is MainCategory Setter
// 门店主营类目
func (r *TaobaoQimenStoreCreateAPIRequest) SetMainCategory(_mainCategory int64) error {
	r._mainCategory = _mainCategory
	r.Set("main_category", _mainCategory)
	return nil
}

// Get MainCategory Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetMainCategory() int64 {
	return r._mainCategory
}

// Set is CompanyName Setter
// 商户名称
func (r *TaobaoQimenStoreCreateAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// Get CompanyName Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetCompanyName() string {
	return r._companyName
}

// Set is EndTime Setter
// 关闭营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is StoreStatus Setter
// 门店状态
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreStatus(_storeStatus string) error {
	r._storeStatus = _storeStatus
	r.Set("store_status", _storeStatus)
	return nil
}

// Get StoreStatus Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreStatus() string {
	return r._storeStatus
}

// Set is StoreDescription Setter
// 商户介绍
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreDescription(_storeDescription string) error {
	r._storeDescription = _storeDescription
	r.Set("store_description", _storeDescription)
	return nil
}

// Get StoreDescription Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreDescription() string {
	return r._storeDescription
}

// Set is Address Setter
// 地址信息
func (r *TaobaoQimenStoreCreateAPIRequest) SetAddress(_address *Address) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetAddress() *Address {
	return r._address
}

// Set is ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoQimenStoreCreateAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetShopId() int64 {
	return r._shopId
}

// Set is StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
	r._storeKeeper = _storeKeeper
	r.Set("store_keeper", _storeKeeper)
	return nil
}

// Get StoreKeeper Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreKeeper() *StoreKeeper {
	return r._storeKeeper
}

// Set is StoreType Setter
// 门店的类型
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreType(_storeType string) error {
	r._storeType = _storeType
	r.Set("store_type", _storeType)
	return nil
}

// Get StoreType Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreType() string {
	return r._storeType
}

// Set is StoreCode Setter
// ERP系统中门店的编码
func (r *TaobaoQimenStoreCreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// Set is Remark Setter
// 备注
func (r *TaobaoQimenStoreCreateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoQimenStoreCreateAPIRequest) GetRemark() string {
	return r._remark
}
