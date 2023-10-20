package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripuppcmembergainAPIRequest 航司权益数据回流 API请求
// alitrip.uppc.member.gain
//
// 航司权益数据回流
type AlitripuppcmembergainAPIRequest struct {
	model.Params
	// 请求唯一标识
	_requestId string
	// 权益截止时间（扩展字段）
	_responseJson string
	// 错误提示
	_errorMsg string
	// 查询成功
	_statusCode int64
}

// NewAlitripuppcmembergainRequest 初始化AlitripuppcmembergainAPIRequest对象
func NewAlitripuppcmembergainRequest() *AlitripuppcmembergainAPIRequest {
	return &AlitripuppcmembergainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripuppcmembergainAPIRequest) GetApiMethodName() string {
	return "alitrip.uppc.member.gain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripuppcmembergainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripuppcmembergainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestId is RequestId Setter
// 请求唯一标识
func (r *AlitripuppcmembergainAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r AlitripuppcmembergainAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetResponseJson is ResponseJson Setter
// 权益截止时间（扩展字段）
func (r *AlitripuppcmembergainAPIRequest) SetResponseJson(_responseJson string) error {
	r._responseJson = _responseJson
	r.Set("response_json", _responseJson)
	return nil
}

// GetResponseJson ResponseJson Getter
func (r AlitripuppcmembergainAPIRequest) GetResponseJson() string {
	return r._responseJson
}

// SetErrorMsg is ErrorMsg Setter
// 错误提示
func (r *AlitripuppcmembergainAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r AlitripuppcmembergainAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetStatusCode is StatusCode Setter
// 查询成功
func (r *AlitripuppcmembergainAPIRequest) SetStatusCode(_statusCode int64) error {
	r._statusCode = _statusCode
	r.Set("status_code", _statusCode)
	return nil
}

// GetStatusCode StatusCode Getter
func (r AlitripuppcmembergainAPIRequest) GetStatusCode() int64 {
	return r._statusCode
}
