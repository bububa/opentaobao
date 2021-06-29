package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询综合体内的门店列表信息 API请求
taobao.koubei.mall.common.store.page

分页查询综合体内的门店列表信息
*/
type TaobaoKoubeiMallCommonStorePageRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string
    // 商圈ID
    mallId   string
    // 分页查询起始值，默认为0
    start   int64
    // 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
    pageSize   int64
    // 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
    categoryIds   []string
    // 商圈内的门店ID
    storeId   string
    // 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
    order   string
    // 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
    serviceTag   []string
    // 口碑城市编码（示例：杭州市330100）
    cityCode   string
    // 经度（终端设备地理位置）
    longitude   string
    // 纬度（终端设备地理位置）
    latitude   string
    // 终端设备描述(中、英文均可)
    terminalType   string
    // 支付宝/口碑/淘宝app版本号
    appVersion   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonStorePageRequest对象
func NewTaobaoKoubeiMallCommonStorePageRequest() *TaobaoKoubeiMallCommonStorePageRequest{
    return &TaobaoKoubeiMallCommonStorePageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStorePageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStorePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetDataSetId() string {
    return r.dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetMallId() string {
    return r.mallId
}
// Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetStart() int64 {
    return r.start
}
// PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetPageSize() int64 {
    return r.pageSize
}
// CategoryIds Setter
// 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetCategoryIds(categoryIds []string) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

// CategoryIds Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetCategoryIds() []string {
    return r.categoryIds
}
// StoreId Setter
// 商圈内的门店ID
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetStoreId() string {
    return r.storeId
}
// Order Setter
// 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

// Order Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetOrder() string {
    return r.order
}
// ServiceTag Setter
// 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetServiceTag(serviceTag []string) error {
    r.serviceTag = serviceTag
    r.Set("service_tag", serviceTag)
    return nil
}

// ServiceTag Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetServiceTag() []string {
    return r.serviceTag
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetCityCode() string {
    return r.cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetLatitude() string {
    return r.latitude
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetTerminalType() string {
    return r.terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetAppVersion() string {
    return r.appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStorePageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStorePageRequest) GetDisplayChannel() string {
    return r.displayChannel
}
