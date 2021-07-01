package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品详情信息 API请求
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息
*/
type TaobaoKoubeiMallCommonItemDetailQueryAPIRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 商圈ID
    _mallId   string
    // 门店ID
    _storeId   string
    // 商品ID
    _itemId   string
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 经度（终端设备地理位置）
    _longitude   string
    // 纬度（终端设备地理位置）
    _latitude   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
}

// 初始化TaobaoKoubeiMallCommonItemDetailQueryAPIRequest对象
func NewTaobaoKoubeiMallCommonItemDetailQueryRequest() *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest{
    return &TaobaoKoubeiMallCommonItemDetailQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetDataSetId() string {
    return r._dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetMallId() string {
    return r._mallId
}
// StoreId Setter
// 门店ID
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetStoreId() string {
    return r._storeId
}
// ItemId Setter
// 商品ID
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetItemId() string {
    return r._itemId
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetCityCode() string {
    return r._cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetLatitude() string {
    return r._latitude
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetDisplayChannel() string {
    return r._displayChannel
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetAppVersion() string {
    return r._appVersion
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonItemDetailQueryAPIRequest) GetTerminalType() string {
    return r._terminalType
}
