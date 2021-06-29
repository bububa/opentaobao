package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店新增接口 API请求
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上
*/
type TaobaoQimenStoreCreateRequest struct {
    model.Params
    // 门店名称
    _storeName   string
    // 门店主营类目
    _mainCategory   int64
    // 商户名称
    _companyName   string
    // 关闭营业时间(只填时，分；只支持半点和整点)
    _endTime   string
    // 开始营业时间(只填时，分；只支持半点和整点)
    _startTime   string
    // 门店状态
    _storeStatus   string
    // 商户介绍
    _storeDescription   string
    // 地址信息
    _address   *Address
    // 需要关联的线上店铺ID
    _shopId   int64
    // 门店所有者信息
    _storeKeeper   *StoreKeeper
    // 门店的类型
    _storeType   string
    // ERP系统中门店的编码
    _storeCode   string
    // 备注
    _remark   string
}

// 初始化TaobaoQimenStoreCreateRequest对象
func NewTaobaoQimenStoreCreateRequest() *TaobaoQimenStoreCreateRequest{
    return &TaobaoQimenStoreCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.store.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreName Setter
// 门店名称
func (r *TaobaoQimenStoreCreateRequest) SetStoreName(_storeName string) error {
    r._storeName = _storeName
    r.Set("store_name", _storeName)
    return nil
}

// StoreName Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreName() string {
    return r._storeName
}
// MainCategory Setter
// 门店主营类目
func (r *TaobaoQimenStoreCreateRequest) SetMainCategory(_mainCategory int64) error {
    r._mainCategory = _mainCategory
    r.Set("main_category", _mainCategory)
    return nil
}

// MainCategory Getter
func (r TaobaoQimenStoreCreateRequest) GetMainCategory() int64 {
    return r._mainCategory
}
// CompanyName Setter
// 商户名称
func (r *TaobaoQimenStoreCreateRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r TaobaoQimenStoreCreateRequest) GetCompanyName() string {
    return r._companyName
}
// EndTime Setter
// 关闭营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoQimenStoreCreateRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoQimenStoreCreateRequest) GetStartTime() string {
    return r._startTime
}
// StoreStatus Setter
// 门店状态
func (r *TaobaoQimenStoreCreateRequest) SetStoreStatus(_storeStatus string) error {
    r._storeStatus = _storeStatus
    r.Set("store_status", _storeStatus)
    return nil
}

// StoreStatus Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreStatus() string {
    return r._storeStatus
}
// StoreDescription Setter
// 商户介绍
func (r *TaobaoQimenStoreCreateRequest) SetStoreDescription(_storeDescription string) error {
    r._storeDescription = _storeDescription
    r.Set("store_description", _storeDescription)
    return nil
}

// StoreDescription Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreDescription() string {
    return r._storeDescription
}
// Address Setter
// 地址信息
func (r *TaobaoQimenStoreCreateRequest) SetAddress(_address *Address) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoQimenStoreCreateRequest) GetAddress() *Address {
    return r._address
}
// ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoQimenStoreCreateRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoQimenStoreCreateRequest) GetShopId() int64 {
    return r._shopId
}
// StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoQimenStoreCreateRequest) SetStoreKeeper(_storeKeeper *StoreKeeper) error {
    r._storeKeeper = _storeKeeper
    r.Set("store_keeper", _storeKeeper)
    return nil
}

// StoreKeeper Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreKeeper() *StoreKeeper {
    return r._storeKeeper
}
// StoreType Setter
// 门店的类型
func (r *TaobaoQimenStoreCreateRequest) SetStoreType(_storeType string) error {
    r._storeType = _storeType
    r.Set("store_type", _storeType)
    return nil
}

// StoreType Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreType() string {
    return r._storeType
}
// StoreCode Setter
// ERP系统中门店的编码
func (r *TaobaoQimenStoreCreateRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreCode() string {
    return r._storeCode
}
// Remark Setter
// 备注
func (r *TaobaoQimenStoreCreateRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenStoreCreateRequest) GetRemark() string {
    return r._remark
}
