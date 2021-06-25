package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店新增接口 APIRequest
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

func NewTaobaoQimenStoreCreateRequest() *TaobaoQimenStoreCreateRequest{
    return &TaobaoQimenStoreCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.store.create"
}

func (r TaobaoQimenStoreCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreCreateRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreName() string {
    return r.storeName
}

func (r *TaobaoQimenStoreCreateRequest) SetMainCategory(mainCategory int64) error {
    r.mainCategory = mainCategory
    r.Set("main_category", mainCategory)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetMainCategory() int64 {
    return r.mainCategory
}

func (r *TaobaoQimenStoreCreateRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetCompanyName() string {
    return r.companyName
}

func (r *TaobaoQimenStoreCreateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoQimenStoreCreateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoQimenStoreCreateRequest) SetStoreStatus(storeStatus string) error {
    r.storeStatus = storeStatus
    r.Set("store_status", storeStatus)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreStatus() string {
    return r.storeStatus
}

func (r *TaobaoQimenStoreCreateRequest) SetStoreDescription(storeDescription string) error {
    r.storeDescription = storeDescription
    r.Set("store_description", storeDescription)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreDescription() string {
    return r.storeDescription
}

func (r *TaobaoQimenStoreCreateRequest) SetAddress(address *Address) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetAddress() *Address {
    return r.address
}

func (r *TaobaoQimenStoreCreateRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetShopId() int64 {
    return r.shopId
}

func (r *TaobaoQimenStoreCreateRequest) SetStoreKeeper(storeKeeper *StoreKeeper) error {
    r.storeKeeper = storeKeeper
    r.Set("store_keeper", storeKeeper)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreKeeper() *StoreKeeper {
    return r.storeKeeper
}

func (r *TaobaoQimenStoreCreateRequest) SetStoreType(storeType string) error {
    r.storeType = storeType
    r.Set("store_type", storeType)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreType() string {
    return r.storeType
}

func (r *TaobaoQimenStoreCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoQimenStoreCreateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoQimenStoreCreateRequest) GetRemark() string {
    return r.remark
}

