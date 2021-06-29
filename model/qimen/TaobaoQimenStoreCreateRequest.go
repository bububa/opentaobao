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
    storeName   string
    // 门店主营类目
    mainCategory   int64
    // 商户名称
    companyName   string
    // 关闭营业时间(只填时，分；只支持半点和整点)
    endTime   string
    // 开始营业时间(只填时，分；只支持半点和整点)
    startTime   string
    // 门店状态
    storeStatus   string
    // 商户介绍
    storeDescription   string
    // 地址信息
    address   *Address
    // 需要关联的线上店铺ID
    shopId   int64
    // 门店所有者信息
    storeKeeper   *StoreKeeper
    // 门店的类型
    storeType   string
    // ERP系统中门店的编码
    storeCode   string
    // 备注
    remark   string
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
func (r *TaobaoQimenStoreCreateRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

// StoreName Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreName() string {
    return r.storeName
}
// MainCategory Setter
// 门店主营类目
func (r *TaobaoQimenStoreCreateRequest) SetMainCategory(mainCategory int64) error {
    r.mainCategory = mainCategory
    r.Set("main_category", mainCategory)
    return nil
}

// MainCategory Getter
func (r TaobaoQimenStoreCreateRequest) GetMainCategory() int64 {
    return r.mainCategory
}
// CompanyName Setter
// 商户名称
func (r *TaobaoQimenStoreCreateRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

// CompanyName Getter
func (r TaobaoQimenStoreCreateRequest) GetCompanyName() string {
    return r.companyName
}
// EndTime Setter
// 关闭营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoQimenStoreCreateRequest) GetEndTime() string {
    return r.endTime
}
// StartTime Setter
// 开始营业时间(只填时，分；只支持半点和整点)
func (r *TaobaoQimenStoreCreateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoQimenStoreCreateRequest) GetStartTime() string {
    return r.startTime
}
// StoreStatus Setter
// 门店状态
func (r *TaobaoQimenStoreCreateRequest) SetStoreStatus(storeStatus string) error {
    r.storeStatus = storeStatus
    r.Set("store_status", storeStatus)
    return nil
}

// StoreStatus Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreStatus() string {
    return r.storeStatus
}
// StoreDescription Setter
// 商户介绍
func (r *TaobaoQimenStoreCreateRequest) SetStoreDescription(storeDescription string) error {
    r.storeDescription = storeDescription
    r.Set("store_description", storeDescription)
    return nil
}

// StoreDescription Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreDescription() string {
    return r.storeDescription
}
// Address Setter
// 地址信息
func (r *TaobaoQimenStoreCreateRequest) SetAddress(address *Address) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoQimenStoreCreateRequest) GetAddress() *Address {
    return r.address
}
// ShopId Setter
// 需要关联的线上店铺ID
func (r *TaobaoQimenStoreCreateRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TaobaoQimenStoreCreateRequest) GetShopId() int64 {
    return r.shopId
}
// StoreKeeper Setter
// 门店所有者信息
func (r *TaobaoQimenStoreCreateRequest) SetStoreKeeper(storeKeeper *StoreKeeper) error {
    r.storeKeeper = storeKeeper
    r.Set("store_keeper", storeKeeper)
    return nil
}

// StoreKeeper Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreKeeper() *StoreKeeper {
    return r.storeKeeper
}
// StoreType Setter
// 门店的类型
func (r *TaobaoQimenStoreCreateRequest) SetStoreType(storeType string) error {
    r.storeType = storeType
    r.Set("store_type", storeType)
    return nil
}

// StoreType Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreType() string {
    return r.storeType
}
// StoreCode Setter
// ERP系统中门店的编码
func (r *TaobaoQimenStoreCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoQimenStoreCreateRequest) GetStoreCode() string {
    return r.storeCode
}
// Remark Setter
// 备注
func (r *TaobaoQimenStoreCreateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenStoreCreateRequest) GetRemark() string {
    return r.remark
}
