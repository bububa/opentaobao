package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonItemShelfPageAPIRequest 门店货架商品列表信息查询 API请求
// taobao.koubei.mall.common.item.shelf.page
//
// 查询口碑综合体内门店货架商品列表信息接口
type TaobaoKoubeiMallCommonItemShelfPageAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 商圈ID
	_mallId string
	// 商圈内的门店ID
	_storeId string
	// 分页查询起始值，默认为0
	_start int64
	// 每页查询量，固定8个
	_pageSize int64
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
}

// NewTaobaoKoubeiMallCommonItemShelfPageRequest 初始化TaobaoKoubeiMallCommonItemShelfPageAPIRequest对象
func NewTaobaoKoubeiMallCommonItemShelfPageRequest() *TaobaoKoubeiMallCommonItemShelfPageAPIRequest {
	return &TaobaoKoubeiMallCommonItemShelfPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.item.shelf.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDataSetId is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetMallId is MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetMallId() string {
	return r._mallId
}

// SetStoreId is StoreId Setter
// 商圈内的门店ID
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetStart is Start Setter
// 分页查询起始值，默认为0
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetStart() int64 {
	return r._start
}

// SetPageSize is PageSize Setter
// 每页查询量，固定8个
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCityCode is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetLongitude is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetDisplayChannel is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// GetDisplayChannel DisplayChannel Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// SetTerminalType is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetAppVersion is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonItemShelfPageAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaoKoubeiMallCommonItemShelfPageAPIRequest) GetAppVersion() string {
	return r._appVersion
}
