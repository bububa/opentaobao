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
type TaobaoKoubeiMallCommonMallAuthPageAPIRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 分页查询起始值，默认为0
    _start   int64
    // 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
    _pageSize   int64
    // 经度（终端设备地理位置）
    _longitude   string
    // 纬度（终端设备地理位置）
    _latitude   string
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonMallAuthPageAPIRequest对象
func NewTaobaoKoubeiMallCommonMallAuthPageRequest() *TaobaoKoubeiMallCommonMallAuthPageAPIRequest{
    return &TaobaoKoubeiMallCommonMallAuthPageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.mall.auth.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetDataSetId() string {
    return r._dataSetId
}
// Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetStart() int64 {
    return r._start
}
// PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetLatitude() string {
    return r._latitude
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetCityCode() string {
    return r._cityCode
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetAppVersion() string {
    return r._appVersion
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonMallAuthPageAPIRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonMallAuthPageAPIRequest) GetDisplayChannel() string {
    return r._displayChannel
}
