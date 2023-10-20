package koubeimall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonMallDetailGetAPIRequest 查询商圈详细信息 API请求
// taobao.koubei.mall.common.mall.detail.get
//
// 查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
type TaobaoKoubeiMallCommonMallDetailGetAPIRequest struct {
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

// NewTaobaoKoubeiMallCommonMallDetailGetRequest 初始化TaobaoKoubeiMallCommonMallDetailGetAPIRequest对象
func NewTaobaoKoubeiMallCommonMallDetailGetRequest() *TaobaoKoubeiMallCommonMallDetailGetAPIRequest {
	return &TaobaoKoubeiMallCommonMallDetailGetAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) Reset() {
	r._cityCode = ""
	r._dataSetId = ""
	r._mallId = ""
	r._displayChannel = ""
	r._appVersion = ""
	r._terminalType = ""
	r._latitude = ""
	r._longitude = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.mall.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetMallId() string {
	return r._mallId
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoKoubeiMallCommonMallDetailGetAPIRequest) GetLongitude() string {
	return r._longitude
}

var poolTaobaoKoubeiMallCommonMallDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoKoubeiMallCommonMallDetailGetRequest()
	},
}

// GetTaobaoKoubeiMallCommonMallDetailGetRequest 从 sync.Pool 获取 TaobaoKoubeiMallCommonMallDetailGetAPIRequest
func GetTaobaoKoubeiMallCommonMallDetailGetAPIRequest() *TaobaoKoubeiMallCommonMallDetailGetAPIRequest {
	return poolTaobaoKoubeiMallCommonMallDetailGetAPIRequest.Get().(*TaobaoKoubeiMallCommonMallDetailGetAPIRequest)
}

// ReleaseTaobaoKoubeiMallCommonMallDetailGetAPIRequest 将 TaobaoKoubeiMallCommonMallDetailGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoKoubeiMallCommonMallDetailGetAPIRequest(v *TaobaoKoubeiMallCommonMallDetailGetAPIRequest) {
	v.Reset()
	poolTaobaoKoubeiMallCommonMallDetailGetAPIRequest.Put(v)
}
