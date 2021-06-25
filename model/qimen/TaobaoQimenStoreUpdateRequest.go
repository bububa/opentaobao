package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店更新接口 APIRequest
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息
*/
type TaobaoQimenStoreUpdateRequest struct {
    model.Params

    // 门店名称
    storeName   string 

    // 备注
    remark   string 

    // 门店主营类目
    mainCategory   int64 

    // 停止营业时间(只填时，分；只支持半点和整点)
    endTime   string 

    // 商户名称
    companyName   string 

    // 开始营业时间(只填时，分；只支持半点和整点)
    startTime   string 

    // 门店状态
    storeStatus   string 

    // 商户介绍
    storeDescription   string 

    // 门店地址信息
    address   *Address 

    // 需要关联的线上店铺ID
    shopId   int64 

    // 门店所有者信息
    storeKeeper   *StoreKeeper 

    // 门店类型
    storeType   string 

    // 线上门店id
    storeId   int64 

    // ERP系统中 门店编码
    storeCode   string 

}

func NewTaobaoQimenStoreUpdateRequest() *TaobaoQimenStoreUpdateRequest{
    return &TaobaoQimenStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreUpdateRequest) GetApiMethodName() string {
    return "taobao.qimen.store.update"
}

func (r TaobaoQimenStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreUpdateRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreName() string {
    return r.storeName
}

func (r *TaobaoQimenStoreUpdateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoQimenStoreUpdateRequest) SetMainCategory(mainCategory int64) error {
    r.mainCategory = mainCategory
    r.Set("main_category", mainCategory)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetMainCategory() int64 {
    return r.mainCategory
}

func (r *TaobaoQimenStoreUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoQimenStoreUpdateRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetCompanyName() string {
    return r.companyName
}

func (r *TaobaoQimenStoreUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreStatus(storeStatus string) error {
    r.storeStatus = storeStatus
    r.Set("store_status", storeStatus)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreStatus() string {
    return r.storeStatus
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreDescription(storeDescription string) error {
    r.storeDescription = storeDescription
    r.Set("store_description", storeDescription)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreDescription() string {
    return r.storeDescription
}

func (r *TaobaoQimenStoreUpdateRequest) SetAddress(address *Address) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetAddress() *Address {
    return r.address
}

func (r *TaobaoQimenStoreUpdateRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetShopId() int64 {
    return r.shopId
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreKeeper(storeKeeper *StoreKeeper) error {
    r.storeKeeper = storeKeeper
    r.Set("store_keeper", storeKeeper)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreKeeper() *StoreKeeper {
    return r.storeKeeper
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreType(storeType string) error {
    r.storeType = storeType
    r.Set("store_type", storeType)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreType() string {
    return r.storeType
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoQimenStoreUpdateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoQimenStoreUpdateRequest) GetStoreCode() string {
    return r.storeCode
}

