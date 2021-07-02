package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthDeliveryBatchaskAPIRequest 广告投放批量询问 API请求
// taobao.usergrowth.delivery.batchask
//
// 提供给媒体在曝光广告前调用， 返回是否曝光以及报价
type TaobaoUsergrowthDeliveryBatchaskAPIRequest struct {
	model.Params
	// 渠道标识，向淘宝技术申请
	_channel string
	// 广告id，淘宝和媒体协商
	_adid string
	// idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
	_idfaMd5 string
	// imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
	_imeiMd5 string
}

// NewTaobaoUsergrowthDeliveryBatchaskRequest 初始化TaobaoUsergrowthDeliveryBatchaskAPIRequest对象
func NewTaobaoUsergrowthDeliveryBatchaskRequest() *TaobaoUsergrowthDeliveryBatchaskAPIRequest {
	return &TaobaoUsergrowthDeliveryBatchaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.delivery.batchask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannel is Channel Setter
// 渠道标识，向淘宝技术申请
func (r *TaobaoUsergrowthDeliveryBatchaskAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetChannel() string {
	return r._channel
}

// SetAdid is Adid Setter
// 广告id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryBatchaskAPIRequest) SetAdid(_adid string) error {
	r._adid = _adid
	r.Set("adid", _adid)
	return nil
}

// GetAdid Adid Getter
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetAdid() string {
	return r._adid
}

// SetIdfaMd5 is IdfaMd5 Setter
// idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
	r._idfaMd5 = _idfaMd5
	r.Set("idfa_md5", _idfaMd5)
	return nil
}

// GetIdfaMd5 IdfaMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetIdfaMd5() string {
	return r._idfaMd5
}

// SetImeiMd5 is ImeiMd5 Setter
// imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDeliveryBatchaskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
	r._imeiMd5 = _imeiMd5
	r.Set("imei_md5", _imeiMd5)
	return nil
}

// GetImeiMd5 ImeiMd5 Getter
func (r TaobaoUsergrowthDeliveryBatchaskAPIRequest) GetImeiMd5() string {
	return r._imeiMd5
}
