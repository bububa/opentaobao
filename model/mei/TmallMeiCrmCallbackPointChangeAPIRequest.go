package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmcallbackpointchangeAPIRequest 品牌积分变更回调API API请求
// tmall.mei.crm.callback.point.change
//
// 线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
type TmallmeicrmcallbackpointchangeAPIRequest struct {
	model.Params
	// 混淆会员手机号码
	_mixMobile string
	// 处理失败的错误码.
	_errorCode string
	// 拓展信息
	_extInfo string
	// 变更记录ID
	_recordId int64
	// 0:成功。1：失败
	_result int64
	// 积分信息
	_point int64
}

// NewTmallmeicrmcallbackpointchangeRequest 初始化TmallmeicrmcallbackpointchangeAPIRequest对象
func NewTmallmeicrmcallbackpointchangeRequest() *TmallmeicrmcallbackpointchangeAPIRequest {
	return &TmallmeicrmcallbackpointchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.callback.point.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixMobile is MixMobile Setter
// 混淆会员手机号码
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetMixMobile(_mixMobile string) error {
	r._mixMobile = _mixMobile
	r.Set("mix_mobile", _mixMobile)
	return nil
}

// GetMixMobile MixMobile Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetMixMobile() string {
	return r._mixMobile
}

// SetErrorCode is ErrorCode Setter
// 处理失败的错误码.
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetExtInfo is ExtInfo Setter
// 拓展信息
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetExtInfo() string {
	return r._extInfo
}

// SetRecordId is RecordId Setter
// 变更记录ID
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetResult is Result Setter
// 0:成功。1：失败
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetResult(_result int64) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetResult() int64 {
	return r._result
}

// SetPoint is Point Setter
// 积分信息
func (r *TmallmeicrmcallbackpointchangeAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// GetPoint Point Getter
func (r TmallmeicrmcallbackpointchangeAPIRequest) GetPoint() int64 {
	return r._point
}
