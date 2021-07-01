package koubeimall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询综合体内的门店列表信息 API请求
taobao.koubei.mall.common.store.page

分页查询综合体内的门店列表信息
*/
type TaobaoKoubeiMallCommonStorePageAPIRequest struct {
    model.Params
    // 身份ID，识别合作方身份（可联系口碑综合体业务获取）
    _dataSetId   string
    // 商圈ID
    _mallId   string
    // 分页查询起始值，默认为0
    _start   int64
    // 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
    _pageSize   int64
    // 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
    _categoryIds   []string
    // 商圈内的门店ID
    _storeId   string
    // 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
    _order   string
    // 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
    _serviceTag   []string
    // 口碑城市编码（示例：杭州市330100）
    _cityCode   string
    // 经度（终端设备地理位置）
    _longitude   string
    // 纬度（终端设备地理位置）
    _latitude   string
    // 终端设备描述(中、英文均可)
    _terminalType   string
    // 支付宝/口碑/淘宝app版本号
    _appVersion   string
    // 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
    _displayChannel   string
}

// 初始化TaobaoKoubeiMallCommonStorePageAPIRequest对象
func NewTaobaoKoubeiMallCommonStorePageRequest() *TaobaoKoubeiMallCommonStorePageAPIRequest{
    return &TaobaoKoubeiMallCommonStorePageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetApiMethodName() string {
    return "taobao.koubei.mall.common.store.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetDataSetId() string {
    return r._dataSetId
}
// MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetMallId() string {
    return r._mallId
}
// Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetStart() int64 {
    return r._start
}
// PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过20）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// CategoryIds Setter
// 门店列表按照类目筛选条件过滤，可通过查询商圈详情获取类目信息
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetCategoryIds(_categoryIds []string) error {
    r._categoryIds = _categoryIds
    r.Set("category_ids", _categoryIds)
    return nil
}

// CategoryIds Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetCategoryIds() []string {
    return r._categoryIds
}
// StoreId Setter
// 商圈内的门店ID
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetStoreId() string {
    return r._storeId
}
// Order Setter
// 门店列表排序规则；默认：门店质量分降序，暂无其它排序规则
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetOrder(_order string) error {
    r._order = _order
    r.Set("order", _order)
    return nil
}

// Order Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetOrder() string {
    return r._order
}
// ServiceTag Setter
// 店铺服务标签，用于列表过滤条件；比如：点餐/外卖/预定等服务筛选条件。预定：SERVICE_DING；排号：SERVICE_PAI；点菜：SERVICE_DIAN；外卖：SERVICE_WAI；
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetServiceTag(_serviceTag []string) error {
    r._serviceTag = _serviceTag
    r.Set("service_tag", _serviceTag)
    return nil
}

// ServiceTag Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetServiceTag() []string {
    return r._serviceTag
}
// CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetCityCode() string {
    return r._cityCode
}
// Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetLatitude() string {
    return r._latitude
}
// TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetAppVersion() string {
    return r._appVersion
}
// DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStorePageAPIRequest) SetDisplayChannel(_displayChannel string) error {
    r._displayChannel = _displayChannel
    r.Set("display_channel", _displayChannel)
    return nil
}

// DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStorePageAPIRequest) GetDisplayChannel() string {
    return r._displayChannel
}
