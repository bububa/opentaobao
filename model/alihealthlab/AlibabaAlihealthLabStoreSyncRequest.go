package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测业务，isv门店同步到健康 APIRequest
alibaba.alihealth.lab.store.sync

阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
*/
type AlibabaAlihealthLabStoreSyncRequest struct {
    model.Params

    // EFFECTIVE 生效，INVALID 失效
    isvStoreStatus   string 

    // 预约须知
    reserveNotice   string 

    // 支持在线报告
    supportOnlineReport   bool 

    // 门店类型描述
    storeTypeDesc   string 

    // 企业社会征信号
    socialCreditCode   string 

    // 营业执照编号
    licenseNo   string 

    // 营业执照名称
    licenseName   string 

    // 门店交通路线
    storeRoutesDesc   string 

    // 营业时间描述
    workTimeDesc   string 

    // 门店电话号码
    storePhone   string 

    // 门店介绍
    storeIntro   string 

    // 经度
    longitude   *BigDecimal 

    // 纬度
    latitude   *BigDecimal 

    // 城市编码
    cityCode   int64 

    // 门店地址
    storeAddress   string 

    // isv门店编码
    isvStoreCode   string 

    // 门店名称
    storeName   string 

    // 支持的淘宝商品类目ID，阿里医院场景
    allowedTbItemCategoryIds   []int64 

}

func NewAlibabaAlihealthLabStoreSyncRequest() *AlibabaAlihealthLabStoreSyncRequest{
    return &AlibabaAlihealthLabStoreSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.store.sync"
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthLabStoreSyncRequest) SetIsvStoreStatus(isvStoreStatus string) error {
    r.isvStoreStatus = isvStoreStatus
    r.Set("isv_store_status", isvStoreStatus)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetIsvStoreStatus() string {
    return r.isvStoreStatus
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetReserveNotice(reserveNotice string) error {
    r.reserveNotice = reserveNotice
    r.Set("reserve_notice", reserveNotice)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetReserveNotice() string {
    return r.reserveNotice
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetSupportOnlineReport(supportOnlineReport bool) error {
    r.supportOnlineReport = supportOnlineReport
    r.Set("support_online_report", supportOnlineReport)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetSupportOnlineReport() bool {
    return r.supportOnlineReport
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreTypeDesc(storeTypeDesc string) error {
    r.storeTypeDesc = storeTypeDesc
    r.Set("store_type_desc", storeTypeDesc)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreTypeDesc() string {
    return r.storeTypeDesc
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetSocialCreditCode(socialCreditCode string) error {
    r.socialCreditCode = socialCreditCode
    r.Set("social_credit_code", socialCreditCode)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetSocialCreditCode() string {
    return r.socialCreditCode
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetLicenseNo(licenseNo string) error {
    r.licenseNo = licenseNo
    r.Set("license_no", licenseNo)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetLicenseNo() string {
    return r.licenseNo
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetLicenseName(licenseName string) error {
    r.licenseName = licenseName
    r.Set("license_name", licenseName)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetLicenseName() string {
    return r.licenseName
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreRoutesDesc(storeRoutesDesc string) error {
    r.storeRoutesDesc = storeRoutesDesc
    r.Set("store_routes_desc", storeRoutesDesc)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreRoutesDesc() string {
    return r.storeRoutesDesc
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetWorkTimeDesc(workTimeDesc string) error {
    r.workTimeDesc = workTimeDesc
    r.Set("work_time_desc", workTimeDesc)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetWorkTimeDesc() string {
    return r.workTimeDesc
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStorePhone(storePhone string) error {
    r.storePhone = storePhone
    r.Set("store_phone", storePhone)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStorePhone() string {
    return r.storePhone
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreIntro(storeIntro string) error {
    r.storeIntro = storeIntro
    r.Set("store_intro", storeIntro)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreIntro() string {
    return r.storeIntro
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetLongitude(longitude *BigDecimal) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetLongitude() *BigDecimal {
    return r.longitude
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetLatitude(latitude *BigDecimal) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetLatitude() *BigDecimal {
    return r.latitude
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetCityCode() int64 {
    return r.cityCode
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreAddress(storeAddress string) error {
    r.storeAddress = storeAddress
    r.Set("store_address", storeAddress)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreAddress() string {
    return r.storeAddress
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetIsvStoreCode(isvStoreCode string) error {
    r.isvStoreCode = isvStoreCode
    r.Set("isv_store_code", isvStoreCode)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetIsvStoreCode() string {
    return r.isvStoreCode
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreName() string {
    return r.storeName
}

func (r *AlibabaAlihealthLabStoreSyncRequest) SetAllowedTbItemCategoryIds(allowedTbItemCategoryIds []int64) error {
    r.allowedTbItemCategoryIds = allowedTbItemCategoryIds
    r.Set("allowed_tb_item_category_ids", allowedTbItemCategoryIds)
    return nil
}

func (r AlibabaAlihealthLabStoreSyncRequest) GetAllowedTbItemCategoryIds() []int64 {
    return r.allowedTbItemCategoryIds
}

