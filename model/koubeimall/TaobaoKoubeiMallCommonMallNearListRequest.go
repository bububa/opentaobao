package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据POI查询附近商圈列表信息 API请求
taobao.koubei.mall.common.mall.near.list

通过用户/终端设备地理位置POI信息，查询附近商圈信息
*/
type TaobaoKoubeiMallCommonMallNearListRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
    _radius   int64
    // 查询个数，最大查询量不能超过50个
    _size   int64
    // 经度（终端设备地理位置）
    _longitude   string
    // 纬度（终端设备地理位置）
    _latitude   string
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
}

// 初始化TaobaoKoubeiMallCommonMallNearListRequest对象
func NewTaobaoKoubeiMallCommonMallNearListRequest() *TaobaoKoubeiMallCommonMallNearListRequest{
    return &TaobaoKoubeiMallCommonMallNearListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.mall.near.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetDataSetId() string {
    return r._dataSetId
}
// Radius Setter
// 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetRadius(_radius int64) error {
    r._radius = _radius
    r.Set("radius", _radius)
    return nil
}

// Radius Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetRadius() int64 {
    return r._radius
}
// Size Setter
// 查询个数，最大查询量不能超过50个
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetSize(_size int64) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetSize() int64 {
    return r._size
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetLatitude() string {
    return r._latitude
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetCityCode() string {
    return r._cityCode
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetDisplayChannel() string {
    return r._displayChannel
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetAppVersion() string {
    return r._appVersion
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonMallNearListRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonMallNearListRequest) GetTerminalType() string {
    return r._terminalType
}
