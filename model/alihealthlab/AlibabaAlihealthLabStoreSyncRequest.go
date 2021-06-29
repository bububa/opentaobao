package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测业务，isv门店同步到健康 API请求
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

// 初始化AlibabaAlihealthLabStoreSyncRequest对象
func NewAlibabaAlihealthLabStoreSyncRequest() *AlibabaAlihealthLabStoreSyncRequest{
    return &AlibabaAlihealthLabStoreSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabStoreSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.store.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabStoreSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvStoreStatus Setter
// EFFECTIVE 生效，INVALID 失效
func (r *AlibabaAlihealthLabStoreSyncRequest) SetIsvStoreStatus(isvStoreStatus string) error {
    r.isvStoreStatus = isvStoreStatus
    r.Set("isv_store_status", isvStoreStatus)
    return nil
}

// IsvStoreStatus Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetIsvStoreStatus() string {
    return r.isvStoreStatus
}
// ReserveNotice Setter
// 预约须知
func (r *AlibabaAlihealthLabStoreSyncRequest) SetReserveNotice(reserveNotice string) error {
    r.reserveNotice = reserveNotice
    r.Set("reserve_notice", reserveNotice)
    return nil
}

// ReserveNotice Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetReserveNotice() string {
    return r.reserveNotice
}
// SupportOnlineReport Setter
// 支持在线报告
func (r *AlibabaAlihealthLabStoreSyncRequest) SetSupportOnlineReport(supportOnlineReport bool) error {
    r.supportOnlineReport = supportOnlineReport
    r.Set("support_online_report", supportOnlineReport)
    return nil
}

// SupportOnlineReport Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetSupportOnlineReport() bool {
    return r.supportOnlineReport
}
// StoreTypeDesc Setter
// 门店类型描述
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreTypeDesc(storeTypeDesc string) error {
    r.storeTypeDesc = storeTypeDesc
    r.Set("store_type_desc", storeTypeDesc)
    return nil
}

// StoreTypeDesc Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreTypeDesc() string {
    return r.storeTypeDesc
}
// SocialCreditCode Setter
// 企业社会征信号
func (r *AlibabaAlihealthLabStoreSyncRequest) SetSocialCreditCode(socialCreditCode string) error {
    r.socialCreditCode = socialCreditCode
    r.Set("social_credit_code", socialCreditCode)
    return nil
}

// SocialCreditCode Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetSocialCreditCode() string {
    return r.socialCreditCode
}
// LicenseNo Setter
// 营业执照编号
func (r *AlibabaAlihealthLabStoreSyncRequest) SetLicenseNo(licenseNo string) error {
    r.licenseNo = licenseNo
    r.Set("license_no", licenseNo)
    return nil
}

// LicenseNo Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetLicenseNo() string {
    return r.licenseNo
}
// LicenseName Setter
// 营业执照名称
func (r *AlibabaAlihealthLabStoreSyncRequest) SetLicenseName(licenseName string) error {
    r.licenseName = licenseName
    r.Set("license_name", licenseName)
    return nil
}

// LicenseName Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetLicenseName() string {
    return r.licenseName
}
// StoreRoutesDesc Setter
// 门店交通路线
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreRoutesDesc(storeRoutesDesc string) error {
    r.storeRoutesDesc = storeRoutesDesc
    r.Set("store_routes_desc", storeRoutesDesc)
    return nil
}

// StoreRoutesDesc Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreRoutesDesc() string {
    return r.storeRoutesDesc
}
// WorkTimeDesc Setter
// 营业时间描述
func (r *AlibabaAlihealthLabStoreSyncRequest) SetWorkTimeDesc(workTimeDesc string) error {
    r.workTimeDesc = workTimeDesc
    r.Set("work_time_desc", workTimeDesc)
    return nil
}

// WorkTimeDesc Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetWorkTimeDesc() string {
    return r.workTimeDesc
}
// StorePhone Setter
// 门店电话号码
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStorePhone(storePhone string) error {
    r.storePhone = storePhone
    r.Set("store_phone", storePhone)
    return nil
}

// StorePhone Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStorePhone() string {
    return r.storePhone
}
// StoreIntro Setter
// 门店介绍
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreIntro(storeIntro string) error {
    r.storeIntro = storeIntro
    r.Set("store_intro", storeIntro)
    return nil
}

// StoreIntro Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreIntro() string {
    return r.storeIntro
}
// Longitude Setter
// 经度
func (r *AlibabaAlihealthLabStoreSyncRequest) SetLongitude(longitude *BigDecimal) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetLongitude() *BigDecimal {
    return r.longitude
}
// Latitude Setter
// 纬度
func (r *AlibabaAlihealthLabStoreSyncRequest) SetLatitude(latitude *BigDecimal) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetLatitude() *BigDecimal {
    return r.latitude
}
// CityCode Setter
// 城市编码
func (r *AlibabaAlihealthLabStoreSyncRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetCityCode() int64 {
    return r.cityCode
}
// StoreAddress Setter
// 门店地址
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreAddress(storeAddress string) error {
    r.storeAddress = storeAddress
    r.Set("store_address", storeAddress)
    return nil
}

// StoreAddress Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreAddress() string {
    return r.storeAddress
}
// IsvStoreCode Setter
// isv门店编码
func (r *AlibabaAlihealthLabStoreSyncRequest) SetIsvStoreCode(isvStoreCode string) error {
    r.isvStoreCode = isvStoreCode
    r.Set("isv_store_code", isvStoreCode)
    return nil
}

// IsvStoreCode Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetIsvStoreCode() string {
    return r.isvStoreCode
}
// StoreName Setter
// 门店名称
func (r *AlibabaAlihealthLabStoreSyncRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

// StoreName Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetStoreName() string {
    return r.storeName
}
// AllowedTbItemCategoryIds Setter
// 支持的淘宝商品类目ID，阿里医院场景
func (r *AlibabaAlihealthLabStoreSyncRequest) SetAllowedTbItemCategoryIds(allowedTbItemCategoryIds []int64) error {
    r.allowedTbItemCategoryIds = allowedTbItemCategoryIds
    r.Set("allowed_tb_item_category_ids", allowedTbItemCategoryIds)
    return nil
}

// AllowedTbItemCategoryIds Getter
func (r AlibabaAlihealthLabStoreSyncRequest) GetAllowedTbItemCategoryIds() []int64 {
    return r.allowedTbItemCategoryIds
}
