package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商圈内的超值特惠商品信息 API请求
taobao.koubei.mall.common.item.super.discount.list

查询商圈超值特惠商品信息列表
*/
type TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 商圈ID
    _mallId   string
    // 查询商品最大个数，最大值50
    _itemSize   int64
    // 经度（终端设备地理位置）
    _longitude   string
    // 纬度（终端设备地理位置）
    _latitude   string
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest对象
func NewTaobaoKoubeiMallCommonItemSuperDiscountListRequest() *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest{
    return &TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.super.discount.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetDataSetId() string {
    return r._dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetMallId() string {
    return r._mallId
}
// ItemSize Setter
// 查询商品最大个数，最大值50
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetItemSize(_itemSize int64) error {
    r._itemSize = _itemSize
    r.Set("item_size", _itemSize)
    return nil
}

// ItemSize Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetItemSize() int64 {
    return r._itemSize
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetLatitude() string {
    return r._latitude
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetCityCode() string {
    return r._cityCode
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetAppVersion() string {
    return r._appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListAPIRequest) GetDisplayChannel() string {
    return r._displayChannel
}
