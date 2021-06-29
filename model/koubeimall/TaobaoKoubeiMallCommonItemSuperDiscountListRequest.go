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
type TaobaoKoubeiMallCommonItemSuperDiscountListRequest struct {
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

// 初始化TaobaoKoubeiMallCommonItemSuperDiscountListRequest对象
func NewTaobaoKoubeiMallCommonItemSuperDiscountListRequest() *TaobaoKoubeiMallCommonItemSuperDiscountListRequest{
    return &TaobaoKoubeiMallCommonItemSuperDiscountListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.item.super.discount.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetDataSetId() string {
    return r._dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetMallId() string {
    return r._mallId
}
// ItemSize Setter
// 查询商品最大个数，最大值50
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetItemSize(_itemSize int64) error {
    r._itemSize = _itemSize
    r.Set("item_size", _itemSize)
    return nil
}

// ItemSize Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetItemSize() int64 {
    return r._itemSize
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetLatitude() string {
    return r._latitude
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetCityCode() string {
    return r._cityCode
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetTerminalType() string {
    return r._terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetAppVersion() string {
    return r._appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonItemSuperDiscountListRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonItemSuperDiscountListRequest) GetDisplayChannel() string {
    return r._displayChannel
}
