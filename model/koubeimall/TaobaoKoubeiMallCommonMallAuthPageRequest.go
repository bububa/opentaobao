package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询已授权的商圈列表信息 API请求
taobao.koubei.mall.common.mall.auth.page

分页查询口碑已授权商圈的列表信息
*/
type TaobaoKoubeiMallCommonMallAuthPageRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string
    // 分页查询起始值，默认为0
    start   int64
    // 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
    pageSize   int64
    // 经度（终端设备地理位置）
    longitude   string
    // 纬度（终端设备地理位置）
    latitude   string
    // 口碑城市编码（示例：杭州市330100）
    cityCode   string
    // 支付宝/口碑/淘宝app版本号
    appVersion   string
    // 终端设备描述(中、英文均可)
    terminalType   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonMallAuthPageRequest对象
func NewTaobaoKoubeiMallCommonMallAuthPageRequest() *TaobaoKoubeiMallCommonMallAuthPageRequest{
    return &TaobaoKoubeiMallCommonMallAuthPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.mall.auth.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetDataSetId() string {
    return r.dataSetId
}
// Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetStart() int64 {
    return r.start
}
// PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetPageSize() int64 {
    return r.pageSize
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetLatitude() string {
    return r.latitude
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetCityCode() string {
    return r.cityCode
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetAppVersion() string {
    return r.appVersion
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetTerminalType() string {
    return r.terminalType
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonMallAuthPageRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonMallAuthPageRequest) GetDisplayChannel() string {
    return r.displayChannel
}
