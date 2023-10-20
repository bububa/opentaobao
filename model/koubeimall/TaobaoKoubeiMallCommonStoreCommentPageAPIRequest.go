package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstorecommentpageAPIRequest 分页查询门店评论详情信息 API请求
// taobao.koubei.mall.common.store.comment.page
//
// 查询口碑综合体内的门店评论信息
type TaobaokoubeimallcommonstorecommentpageAPIRequest struct {
	model.Params
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 经度（终端设备地理位置）
	_latitude string
	// 纬度（终端设备地理位置）
	_longitude string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 门店ID
	_storeId string
	// 每页查询量，默认为20，最大数值20
	_pageSize int64
	// 查询起始值，默认为0
	_start int64
}

// NewTaobaokoubeimallcommonstorecommentpageRequest 初始化TaobaokoubeimallcommonstorecommentpageAPIRequest对象
func NewTaobaokoubeimallcommonstorecommentpageRequest() *TaobaokoubeimallcommonstorecommentpageAPIRequest {
	return &TaobaokoubeimallcommonstorecommentpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.store.comment.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLatitude is Latitude Setter
// 经度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetMallId() string {
	return r._mallId
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetPageSize is PageSize Setter
// 每页查询量，默认为20，最大数值20
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStart is Start Setter
// 查询起始值，默认为0
func (r *TaobaokoubeimallcommonstorecommentpageAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaokoubeimallcommonstorecommentpageAPIRequest) GetStart() int64 {
	return r._start
}
