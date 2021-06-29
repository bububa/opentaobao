package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店更新接口 API请求
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息
*/
type TaobaoQimenStoreUpdateRequest struct {
    model.Params
    // 门店名称
    _storeName   string
    // 备注
    _remark   string
    // 门店主营类目
    _mainCategory   int64
    // 停止营业时间(只填时，分；只支持半点和整点)
    _endTime   string
    // 商户名称
    _companyName   string
    // 开始营业时间(只填时，分；只支持半点和整点)
    _startTime   string
    // 门店状态
    _storeStatus   string
    // 商户介绍
    _storeDescription   string
    // 门店地址信息
    _address   *Address
    // 需要关联的线上店铺ID
    _shopId   int64
    // 门店所有者信息
    _storeKeeper   *StoreKeeper
    // 门店类型
    _storeType   string
    // 线上门店id
    _storeId   int64
    // ERP系统中 门店编码
    _storeCode   string
}

// 初始化TaobaoQimenStoreUpdateRequest对象
func NewTaobaoQimenStoreUpdateRequest() *TaobaoQimenStoreUpdateRequest{
    return &TaobaoQimenStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreUpdateRequest) GetApiMethodName() string {
    return "taobao.qimen.store.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreName Setter
// 门店名称
func (r *TaobaoQimenStoreUpdateRequest) SetStoreName(_storeName string) error {
    r._storeName = _storeName
    r.Set("store_name", _storeName)
    return nil
}

// StoreName Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreName() string {
    return r._storeName
}
// Remark Setter
// 备注
func (r *TaobaoQimenStoreUpdateRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenStoreUpdateRequest) GetRemark() string {
    return r._remark
}
// MainCategory Setter
// 门店主营类目
func (r *TaobaoQimenStoreUpdateRequest) SetMainCategory(_mainCategory int64) error {
    r._mainCategory = _mainCategory
    r.Set("main_category", _mainCategory)
    return nil
}

// MainCategory Getter
func (r TaobaoQimenStoreUpdateRequest) GetMainCategory() int64 {
    return r._mainCategory
}
// EndTime Setter
// 停止营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreUpdateRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoQimenStoreUpdateRequest) GetEndTime() string {
    return r._endTime
}
// CompanyName Setter
// 商户名称
func (r *TaobaoQimenStoreUpdateRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r TaobaoQimenStoreUpdateRequest) GetCompanyName() string {
    return r._companyName
}
// StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreUpdateRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoQimenStoreUpdateRequest) GetStartTime() string {
    return r._startTime
}
// StoreStatus Setter
// 门店状态
func (r *TaobaoQimenStoreUpdateRequest) SetStoreStatus(_storeStatus string) error {
    r._storeStatus = _storeStatus
    r.Set("store_status", _storeStatus)
    return nil
}

// StoreStatus Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreStatus() string {
    return r._storeStatus
}
// StoreDescription Setter
// 商户介绍
func (r *TaobaoQimenStoreUpdateRequest) SetStoreDescription(_storeDescription string) error {
    r._storeDescription = _storeDescription
    r.Set("store_description", _storeDescription)
    return nil
}

// StoreDescription Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreDescription() string {
    return r._storeDescription
}
// Address Setter
// 门店地址信息
func (r *TaobaoQimenStoreUpdateRequest) SetAddress(_address *Address) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoQimenStoreUpdateRequest) GetAddress() *Address {
    return r._address
}
// ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoQimenStoreUpdateRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoQimenStoreUpdateRequest) GetShopId() int64 {
    return r._shopId
}
// StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoQimenStoreUpdateRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
    r._storeKeeper = _storeKeeper
    r.Set("store_keeper", _storeKeeper)
    return nil
}

// StoreKeeper Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreKeeper() *StoreKeeper {
    return r._storeKeeper
}
// StoreType Setter
// 门店类型
func (r *TaobaoQimenStoreUpdateRequest) SetStoreType(_storeType string) error {
    r._storeType = _storeType
    r.Set("store_type", _storeType)
    return nil
}

// StoreType Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreType() string {
    return r._storeType
}
// StoreId Setter
// 线上门店id
func (r *TaobaoQimenStoreUpdateRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreId() int64 {
    return r._storeId
}
// StoreCode Setter
// ERP系统中 门店编码
func (r *TaobaoQimenStoreUpdateRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoQimenStoreUpdateRequest) GetStoreCode() string {
    return r._storeCode
}
