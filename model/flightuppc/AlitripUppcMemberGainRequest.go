package flightuppc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
航司权益数据回流 API请求
alitrip.uppc.member.gain

航司权益数据回流
*/
type AlitripUppcMemberGainRequest struct {
    model.Params
    // 请求唯一标识
    _requestId   string
    // 查询成功
    _statusCode   int64
    // 权益截止时间（扩展字段）
    _responseJson   string
    // 错误提示
    _errorMsg   string
}

// 初始化AlitripUppcMemberGainRequest对象
func NewAlitripUppcMemberGainRequest() *AlitripUppcMemberGainRequest{
    return &AlitripUppcMemberGainRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripUppcMemberGainRequest) GetApiMethodName() string {
    return "alitrip.uppc.member.gain"
}

// IRequest interface 方法, 获取API参数
func (r AlitripUppcMemberGainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestId Setter
// 请求唯一标识
func (r *AlitripUppcMemberGainRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r AlitripUppcMemberGainRequest) GetRequestId() string {
    return r._requestId
}
// StatusCode Setter
// 查询成功
func (r *AlitripUppcMemberGainRequest) SetStatusCode(_statusCode int64) error {
    r._statusCode = _statusCode
    r.Set("status_code", _statusCode)
    return nil
}

// StatusCode Getter
func (r AlitripUppcMemberGainRequest) GetStatusCode() int64 {
    return r._statusCode
}
// ResponseJson Setter
// 权益截止时间（扩展字段）
func (r *AlitripUppcMemberGainRequest) SetResponseJson(_responseJson string) error {
    r._responseJson = _responseJson
    r.Set("response_json", _responseJson)
    return nil
}

// ResponseJson Getter
func (r AlitripUppcMemberGainRequest) GetResponseJson() string {
    return r._responseJson
}
// ErrorMsg Setter
// 错误提示
func (r *AlitripUppcMemberGainRequest) SetErrorMsg(_errorMsg string) error {
    r._errorMsg = _errorMsg
    r.Set("error_msg", _errorMsg)
    return nil
}

// ErrorMsg Getter
func (r AlitripUppcMemberGainRequest) GetErrorMsg() string {
    return r._errorMsg
}
