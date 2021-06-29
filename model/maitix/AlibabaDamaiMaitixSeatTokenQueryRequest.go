package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商选座获取qtoken API请求
alibaba.damai.maitix.seat.token.query

选座分销，分销商查询token
*/
type AlibabaDamaiMaitixSeatTokenQueryRequest struct {
    model.Params
    // 场次ID-必填
    _performId   int64
    // 项目ID-必填
    _projectId   int64
    // 必填-选座结束跳转回去的url,这是渠道方自己的url地址,用于接收选座后的座位信息参数
    _callbackUrl   string
    // 会话ID，保证一次选座会话,建议使用 appKey+随机串 生成 ；注意：同一个场次下的会话ID不能重复
    _requestId   string
}

// 初始化AlibabaDamaiMaitixSeatTokenQueryRequest对象
func NewAlibabaDamaiMaitixSeatTokenQueryRequest() *AlibabaDamaiMaitixSeatTokenQueryRequest{
    return &AlibabaDamaiMaitixSeatTokenQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.seat.token.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PerformId Setter
// 场次ID-必填
func (r *AlibabaDamaiMaitixSeatTokenQueryRequest) SetPerformId(_performId int64) error {
    r._performId = _performId
    r.Set("perform_id", _performId)
    return nil
}

// PerformId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetPerformId() int64 {
    return r._performId
}
// ProjectId Setter
// 项目ID-必填
func (r *AlibabaDamaiMaitixSeatTokenQueryRequest) SetProjectId(_projectId int64) error {
    r._projectId = _projectId
    r.Set("project_id", _projectId)
    return nil
}

// ProjectId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetProjectId() int64 {
    return r._projectId
}
// CallbackUrl Setter
// 必填-选座结束跳转回去的url,这是渠道方自己的url地址,用于接收选座后的座位信息参数
func (r *AlibabaDamaiMaitixSeatTokenQueryRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// RequestId Setter
// 会话ID，保证一次选座会话,建议使用 appKey+随机串 生成 ；注意：同一个场次下的会话ID不能重复
func (r *AlibabaDamaiMaitixSeatTokenQueryRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryRequest) GetRequestId() string {
    return r._requestId
}
