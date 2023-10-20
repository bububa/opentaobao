package usergrowth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest 广告曝光前判定批量接口V2 API请求
// taobao.usergrowth.dhh.delivery.batchask
//
// 广告曝光前判定批量接口V2
type TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest struct {
	model.Params
	// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
	_oaidMd5 string
	// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
	_idfaMd5 string
	// md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
	_imeiMd5 string
	// 巨浪广告位id,在巨浪平台申请
	_advertisingSpaceId string
	// 巨浪渠道id,在巨浪平台申请
	_channel string
	// md5加密后的caid列表， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本， 多个使用,分隔， 最多支持20个。
	_caidMd5 string
}

// NewTaobaoUsergrowthDhhDeliveryBatchaskRequest 初始化TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest对象
func NewTaobaoUsergrowthDhhDeliveryBatchaskRequest() *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest {
	return &TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) Reset() {
	r._oaidMd5 = ""
	r._idfaMd5 = ""
	r._imeiMd5 = ""
	r._advertisingSpaceId = ""
	r._channel = ""
	r._caidMd5 = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.dhh.delivery.batchask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOaidMd5 is OaidMd5 Setter
// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetOaidMd5(_oaidMd5 string) error {
	r._oaidMd5 = _oaidMd5
	r.Set("oaid_md5", _oaidMd5)
	return nil
}

// GetOaidMd5 OaidMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetOaidMd5() string {
	return r._oaidMd5
}

// SetIdfaMd5 is IdfaMd5 Setter
// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
	r._idfaMd5 = _idfaMd5
	r.Set("idfa_md5", _idfaMd5)
	return nil
}

// GetIdfaMd5 IdfaMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetIdfaMd5() string {
	return r._idfaMd5
}

// SetImeiMd5 is ImeiMd5 Setter
// md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
	r._imeiMd5 = _imeiMd5
	r.Set("imei_md5", _imeiMd5)
	return nil
}

// GetImeiMd5 ImeiMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetImeiMd5() string {
	return r._imeiMd5
}

// SetAdvertisingSpaceId is AdvertisingSpaceId Setter
// 巨浪广告位id,在巨浪平台申请
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetAdvertisingSpaceId(_advertisingSpaceId string) error {
	r._advertisingSpaceId = _advertisingSpaceId
	r.Set("advertising_space_id", _advertisingSpaceId)
	return nil
}

// GetAdvertisingSpaceId AdvertisingSpaceId Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetAdvertisingSpaceId() string {
	return r._advertisingSpaceId
}

// SetChannel is Channel Setter
// 巨浪渠道id,在巨浪平台申请
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetChannel() string {
	return r._channel
}

// SetCaidMd5 is CaidMd5 Setter
// md5加密后的caid列表， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本， 多个使用,分隔， 最多支持20个。
func (r *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) SetCaidMd5(_caidMd5 string) error {
	r._caidMd5 = _caidMd5
	r.Set("caid_md5", _caidMd5)
	return nil
}

// GetCaidMd5 CaidMd5 Getter
func (r TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) GetCaidMd5() string {
	return r._caidMd5
}

var poolTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsergrowthDhhDeliveryBatchaskRequest()
	},
}

// GetTaobaoUsergrowthDhhDeliveryBatchaskRequest 从 sync.Pool 获取 TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest
func GetTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest() *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest {
	return poolTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest.Get().(*TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest)
}

// ReleaseTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest 将 TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest(v *TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest) {
	v.Reset()
	poolTaobaoUsergrowthDhhDeliveryBatchaskAPIRequest.Put(v)
}
