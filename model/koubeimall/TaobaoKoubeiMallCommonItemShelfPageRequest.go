package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店货架商品列表信息查询 API请求
taobao.koubei.mall.common.item.shelf.page

查询口碑综合体内门店货架商品列表信息接口
*/
type TaobaoKoubeiMallCommonItemShelfPageRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string
    // 商圈ID
    mallId   string
    // 商圈内的门店ID
    storeId   string
    // 分页查询起始值，默认为0
    start   int64
    // 每页查询量，固定8个
    pageSize   int64
    // 口碑城市编码（示例：杭州市330100）
    cityCode   string
    // 经度（终端设备地理位置）
    longitude   string
    // 纬度（终端设备地理位置）
    latitude   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string
    // 终端设备描述(中、英文均可)
    terminalType   string
    // 支付宝/口碑/淘宝app版本号
    appVersion   string
}

// 初始化TaobaoKoubeiMallCommonItemShelfPageRequest对象
func NewTaobaoKoubeiMallCommonItemShelfPageRequest() *TaobaoKoubeiMallCommonItemShelfPageRequest{
    return &TaobaoKoubeiMallCommonItemShelfPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.shelf.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetDataSetId() string {
    return r.dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetMallId() string {
    return r.mallId
}
// StoreId Setter
// 商圈内的门店ID
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetStoreId() string {
    return r.storeId
}
// Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetStart() int64 {
    return r.start
}
// PageSize Setter
// 每页查询量，固定8个
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetPageSize() int64 {
    return r.pageSize
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetCityCode() string {
    return r.cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetLatitude() string {
    return r.latitude
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetDisplayChannel() string {
    return r.displayChannel
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetTerminalType() string {
    return r.terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonItemShelfPageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonItemShelfPageRequest) GetAppVersion() string {
    return r.appVersion
}
