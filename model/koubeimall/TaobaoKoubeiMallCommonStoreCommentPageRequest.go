package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询门店评论详情信息 API请求
taobao.koubei.mall.common.store.comment.page

查询口碑综合体内的门店评论信息
*/
type TaobaoKoubeiMallCommonStoreCommentPageRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 商圈ID
    _mallId   string
    // 门店ID
    _storeId   string
    // 查询起始值，默认为0
    _start   int64
    // 每页查询量，默认为20，最大数值20
    _pageSize   int64
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 纬度（终端设备地理位置）
    _longitude   string
    // 经度（终端设备地理位置）
    _latitude   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonStoreCommentPageRequest对象
func NewTaobaoKoubeiMallCommonStoreCommentPageRequest() *TaobaoKoubeiMallCommonStoreCommentPageRequest{
    return &TaobaoKoubeiMallCommonStoreCommentPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.comment.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetDataSetId() string {
    return r._dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetMallId() string {
    return r._mallId
}
// StoreId Setter
// 门店ID
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetStoreId() string {
    return r._storeId
}
// Start Setter
// 查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetStart() int64 {
    return r._start
}
// PageSize Setter
// 每页查询量，默认为20，最大数值20
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetPageSize() int64 {
    return r._pageSize
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetCityCode() string {
    return r._cityCode
}
// Longitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetLatitude() string {
    return r._latitude
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetTerminalType() string {
    return r._terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetAppVersion() string {
    return r._appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStoreCommentPageRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStoreCommentPageRequest) GetDisplayChannel() string {
    return r._displayChannel
}
