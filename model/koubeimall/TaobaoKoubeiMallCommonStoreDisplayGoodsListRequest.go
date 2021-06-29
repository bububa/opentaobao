package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店推荐菜信息 API请求
taobao.koubei.mall.common.store.display.goods.list

提供查询口碑商圈内的门店推荐菜信息
*/
type TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    dataSetId   string
    // 门店ID
    storeId   string
    // 商圈ID
    mallId   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    displayChannel   string
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
}

// 初始化TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest对象
func NewTaobaoKoubeiMallCommonStoreDisplayGoodsListRequest() *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest{
    return &TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.display.goods.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetDataSetId() string {
    return r.dataSetId
}
// StoreId Setter
// 门店ID
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetStoreId() string {
    return r.storeId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetMallId() string {
    return r.mallId
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetDisplayChannel(displayChannel string) error {
    r.displayChannel = displayChannel
    r.Set("display_channel", displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetDisplayChannel() string {
    return r.displayChannel
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetCityCode() string {
    return r.cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetLatitude() string {
    return r.latitude
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetTerminalType() string {
    return r.terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest) GetAppVersion() string {
    return r.appVersion
}
