package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonmalldetailgetAPIRequest 查询商圈详细信息 API请求
// taobao.koubei.mall.common.mall.detail.get
//
// 查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
type TaobaokoubeimallcommonmalldetailgetAPIRequest struct {
	model.Params
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
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
}

// NewTaobaokoubeimallcommonmalldetailgetRequest 初始化TaobaokoubeimallcommonmalldetailgetAPIRequest对象
func NewTaobaokoubeimallcommonmalldetailgetRequest() *TaobaokoubeimallcommonmalldetailgetAPIRequest {
	return &TaobaokoubeimallcommonmalldetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.mall.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetMallId() string {
	return r._mallId
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonmalldetailgetAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonmalldetailgetAPIRequest) GetLongitude() string {
	return r._longitude
}
