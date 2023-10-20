package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonmallauthpageAPIRequest 分页查询已授权的商圈列表信息 API请求
// taobao.koubei.mall.common.mall.auth.page
//
// 分页查询口碑已授权商圈的列表信息
type TaobaokoubeimallcommonmallauthpageAPIRequest struct {
	model.Params
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 纬度（终端设备地理位置）
	_latitude string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 经度（终端设备地理位置）
	_longitude string
	// 分页查询起始值，默认为0
	_start int64
	// 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
	_pageSize int64
}

// NewTaobaokoubeimallcommonmallauthpageRequest 初始化TaobaokoubeimallcommonmallauthpageAPIRequest对象
func NewTaobaokoubeimallcommonmallauthpageRequest() *TaobaokoubeimallcommonmallauthpageAPIRequest {
	return &TaobaokoubeimallcommonmallauthpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.mall.auth.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetStart is Start Setter
// 分页查询起始值，默认为0
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetStart() int64 {
	return r._start
}

// SetPageSize is PageSize Setter
// 每页查询量，默认10（建议查询值为10的倍数，最大不超过30）
func (r *TaobaokoubeimallcommonmallauthpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaokoubeimallcommonmallauthpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
