package flightuppc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
航司权益数据回流 APIRequest
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

func NewAlitripUppcMemberGainRequest() *AlitripUppcMemberGainRequest{
    return &AlitripUppcMemberGainRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripUppcMemberGainRequest) GetApiMethodName() string {
    return "alitrip.uppc.member.gain"
}

func (r AlitripUppcMemberGainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripUppcMemberGainRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r AlitripUppcMemberGainRequest) GetRequestId() string {
    return r.requestId
}

func (r *AlitripUppcMemberGainRequest) SetStatusCode(statusCode int64) error {
    r.statusCode = statusCode
    r.Set("status_code", statusCode)
    return nil
}

func (r AlitripUppcMemberGainRequest) GetStatusCode() int64 {
    return r.statusCode
}

func (r *AlitripUppcMemberGainRequest) SetResponseJson(responseJson string) error {
    r.responseJson = responseJson
    r.Set("response_json", responseJson)
    return nil
}

func (r AlitripUppcMemberGainRequest) GetResponseJson() string {
    return r.responseJson
}

func (r *AlitripUppcMemberGainRequest) SetErrorMsg(errorMsg string) error {
    r.errorMsg = errorMsg
    r.Set("error_msg", errorMsg)
    return nil
}

func (r AlitripUppcMemberGainRequest) GetErrorMsg() string {
    return r.errorMsg
}

