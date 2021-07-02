package koubeimall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest 查询门店推荐菜信息 API请求
// taobao.koubei.mall.common.store.display.goods.list
//
// 提供查询口碑商圈内的门店推荐菜信息
type TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest struct {
	model.Params
	// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
	_dataSetId string
	// 门店ID
	_storeId string
	// 商圈ID
	_mallId string
	// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
	_displayChannel string
	// 口碑城市编码（示例：杭州市330100）
	_cityCode string
	// 经度（终端设备地理位置）
	_longitude string
	// 纬度（终端设备地理位置）
	_latitude string
	// 终端设备描述(中、英文均可)
	_terminalType string
	// 支付宝/口碑/淘宝app版本号
	_appVersion string
}

// NewTaobaoKoubeiMallCommonStoreDisplayGoodsListRequest 初始化TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest对象
func NewTaobaoKoubeiMallCommonStoreDisplayGoodsListRequest() *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest {
	return &TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.mall.common.store.display.goods.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DataSetId Setter
// 身份ID，识别合作方身份（可联系口碑综合体业务获取）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// Get DataSetId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is MallId Setter
// 商圈ID
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// Get MallId Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetMallId() string {
	return r._mallId
}

// Set is DisplayChannel Setter
// 展示渠道：ALIPAY_APP KOUBEI_APP TAOBAO_APP（默认ALIPAY_APP）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetDisplayChannel(_displayChannel string) error {
	r._displayChannel = _displayChannel
	r.Set("display_channel", _displayChannel)
	return nil
}

// Get DisplayChannel Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetDisplayChannel() string {
	return r._displayChannel
}

// Set is CityCode Setter
// 口碑城市编码（示例：杭州市330100）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// Get CityCode Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetCityCode() string {
	return r._cityCode
}

// Set is Longitude Setter
// 经度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// Get Longitude Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetLongitude() string {
	return r._longitude
}

// Set is Latitude Setter
// 纬度（终端设备地理位置）
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// Get Latitude Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetLatitude() string {
	return r._latitude
}

// Set is TerminalType Setter
// 终端设备描述(中、英文均可)
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// Get TerminalType Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// Set is AppVersion Setter
// 支付宝/口碑/淘宝app版本号
func (r *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// Get AppVersion Getter
func (r TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest) GetAppVersion() string {
	return r._appVersion
}
