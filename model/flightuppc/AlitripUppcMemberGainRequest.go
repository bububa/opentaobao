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
    requestId   string
    // 查询成功
    statusCode   int64
    // 权益截止时间（扩展字段）
    responseJson   string
    // 错误提示
    errorMsg   string
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
func (r *AlitripUppcMemberGainRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r AlitripUppcMemberGainRequest) GetRequestId() string {
    return r.requestId
}
// StatusCode Setter
// 查询成功
func (r *AlitripUppcMemberGainRequest) SetStatusCode(statusCode int64) error {
    r.statusCode = statusCode
    r.Set("status_code", statusCode)
    return nil
}

// StatusCode Getter
func (r AlitripUppcMemberGainRequest) GetStatusCode() int64 {
    return r.statusCode
}
// ResponseJson Setter
// 权益截止时间（扩展字段）
func (r *AlitripUppcMemberGainRequest) SetResponseJson(responseJson string) error {
    r.responseJson = responseJson
    r.Set("response_json", responseJson)
    return nil
}

// ResponseJson Getter
func (r AlitripUppcMemberGainRequest) GetResponseJson() string {
    return r.responseJson
}
// ErrorMsg Setter
// 错误提示
func (r *AlitripUppcMemberGainRequest) SetErrorMsg(errorMsg string) error {
    r.errorMsg = errorMsg
    r.Set("error_msg", errorMsg)
    return nil
}

// ErrorMsg Getter
func (r AlitripUppcMemberGainRequest) GetErrorMsg() string {
    return r.errorMsg
}
