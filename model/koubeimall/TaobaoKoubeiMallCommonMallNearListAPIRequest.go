package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonmallnearlistAPIRequest 根据POI查询附近商圈列表信息 API请求
// taobao.koubei.mall.common.mall.near.list
//
// 通过用户/终端设备地理位置POI信息，查询附近商圈信息
type TaobaokoubeimallcommonmallnearlistAPIRequest struct {
	model.Params
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
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
	// 查询个数，最大查询量不能超过50个
	_size int64
	// 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
	_radius int64
}

// NewTaobaokoubeimallcommonmallnearlistRequest 初始化TaobaokoubeimallcommonmallnearlistAPIRequest对象
func NewTaobaokoubeimallcommonmallnearlistRequest() *TaobaokoubeimallcommonmallnearlistAPIRequest {
	return &TaobaokoubeimallcommonmallnearlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.mall.near.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetSize is Size Setter
// 查询个数，最大查询量不能超过50个
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetSize() int64 {
	return r._size
}

// SetRadius is Radius Setter
// 召回半径，单位m，最大数值不能超过10km（该字段为空，默认全城召回）
func (r *TaobaokoubeimallcommonmallnearlistAPIRequest) SetRadius(_radius int64) error {
	r._radius = _radius
	r.Set("radius", _radius)
	return nil
}

// GetRadius Radius Getter
func (r TaobaokoubeimallcommonmallnearlistAPIRequest) GetRadius() int64 {
	return r._radius
}
