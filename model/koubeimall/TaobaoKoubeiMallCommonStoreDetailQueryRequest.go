package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询综合体内的门店详细信息 API请求
taobao.koubei.mall.common.store.detail.query

查询口碑综合体内的门店详情信息
*/
type TaobaoKoubeiMallCommonStoreDetailQueryRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string
    // 商圈ID
    mallId   string
    // 商圈内的门店ID
    storeId   string
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

// 初始化TaobaoKoubeiMallCommonStoreDetailQueryRequest对象
func NewTaobaoKoubeiMallCommonStoreDetailQueryRequest() *TaobaoKoubeiMallCommonStoreDetailQueryRequest{
    return &TaobaoKoubeiMallCommonStoreDetailQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetDataSetId() string {
    return r.dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetMallId() string {
    return r.mallId
}
// StoreId Setter
// 商圈内的门店ID
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetStoreId() string {
    return r.storeId
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetCityCode() string {
    return r.cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetLatitude() string {
    return r.latitude
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetTerminalType() string {
    return r.terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetAppVersion() string {
    return r.appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStoreDetailQueryRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStoreDetailQueryRequest) GetDisplayChannel() string {
    return r.displayChannel
}
