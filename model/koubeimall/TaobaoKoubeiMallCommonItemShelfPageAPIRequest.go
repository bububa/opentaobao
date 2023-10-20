package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonitemshelfpageAPIRequest 门店货架商品列表信息查询 API请求
// taobao.koubei.mall.common.item.shelf.page
//
// 查询口碑综合体内门店货架商品列表信息接口
type TaobaokoubeimallcommonitemshelfpageAPIRequest struct {
	model.Params
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 纬度（终端设备地理位置）
	_latitude string
	// 经度（终端设备地理位置）
	_longitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈内的门店ID
	_storeId string
	// 商圈ID
	_mallId string
	// 每页查询量，固定8个
	_pageSize int64
	// 分页查询起始值，默认为0
	_start int64
}

// NewTaobaokoubeimallcommonitemshelfpageRequest 初始化TaobaokoubeimallcommonitemshelfpageAPIRequest对象
func NewTaobaokoubeimallcommonitemshelfpageRequest() *TaobaokoubeimallcommonitemshelfpageAPIRequest {
	return &TaobaokoubeimallcommonitemshelfpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.item.shelf.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetStoreId is StoreId Setter
// 商圈内的门店ID
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetMallId() string {
	return r._mallId
}

// SetPageSize is PageSize Setter
// 每页查询量，固定8个
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStart is Start Setter
// 分页查询起始值，默认为0
func (r *TaobaokoubeimallcommonitemshelfpageAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaokoubeimallcommonitemshelfpageAPIRequest) GetStart() int64 {
	return r._start
}
