package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMeiCrmCallbackPointChangeAPIRequest
品牌积分变更回调API API请求
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。 */
type TmallMeiCrmCallbackPointChangeAPIRequest struct {
	model.Params
	// 混淆会员手机号码
	_mixMobile string
	// 变更记录ID
	_recordId int64
	// 0:成功。1：失败
	_result int64
	// 处理失败的错误码.
	_errorCode string
	// 拓展信息
	_extInfo string
	// 积分信息
	_point int64
}

// NewTmallMeiCrmCallbackPointChangeRequest 初始化TmallMeiCrmCallbackPointChangeAPIRequest对象
func NewTmallMeiCrmCallbackPointChangeRequest() *TmallMeiCrmCallbackPointChangeAPIRequest {
	return &TmallMeiCrmCallbackPointChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.callback.point.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MixMobile Setter
// 混淆会员手机号码
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetMixMobile(_mixMobile string) error {
	r._mixMobile = _mixMobile
	r.Set("mix_mobile", _mixMobile)
	return nil
}

// Get MixMobile Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetMixMobile() string {
	return r._mixMobile
}

// Set is RecordId Setter
// 变更记录ID
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// Get RecordId Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// Set is Result Setter
// 0:成功。1：失败
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetResult(_result int64) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// Get Result Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetResult() int64 {
	return r._result
}

// Set is ErrorCode Setter
// 处理失败的错误码.
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// Get ErrorCode Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// Set is ExtInfo Setter
// 拓展信息
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// Get ExtInfo Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetExtInfo() string {
	return r._extInfo
}

// Set is Point Setter
// 积分信息
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetPoint(_point int64) error {
	r._point = _point
	r.Set("point", _point)
	return nil
}

// Get Point Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetPoint() int64 {
	return r._point
}
