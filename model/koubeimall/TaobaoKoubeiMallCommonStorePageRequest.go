package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询综合体内的门店列表信息 APIRequest
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

func NewTaobaoKoubeiMallCommonStorePageRequest() *TaobaoKoubeiMallCommonStorePageRequest{
    return &TaobaoKoubeiMallCommonStorePageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.page"
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiMallCommonStorePageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetMallId() string {
    return r.mallId
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetStart() int64 {
    return r.start
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetCategoryIds(categoryIds []string) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetCategoryIds() []string {
    return r.categoryIds
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetStoreId() string {
    return r.storeId
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetOrder(order string) error {
    r.order = order
    r.Set("order", order)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetOrder() string {
    return r.order
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetServiceTag(serviceTag []string) error {
    r.serviceTag = serviceTag
    r.Set("service_tag", serviceTag)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetServiceTag() []string {
    return r.serviceTag
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoKoubeiMallCommonStorePageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

func (r TaobaoKoubeiMallCommonStorePageRequest) GetDisplayChannel() string {
    return r.displayChannel
}

