package ascpqcc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品控中心更新样品信息 API请求
alibaba.ascp.qcc.sample.update

品控中心更新样品信息
*/
type AlibabaAscpQccSampleUpdateAPIRequest struct {
    model.Params
    // 更新请求参数
    _updateRequest   *UpdateSampleRequest
}

// 初始化AlibabaAscpQccSampleUpdateAPIRequest对象
func NewAlibabaAscpQccSampleUpdateRequest() *AlibabaAscpQccSampleUpdateAPIRequest{
    return &AlibabaAscpQccSampleUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.qcc.sample.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRequest Setter
// 更新请求参数
func (r *AlibabaAscpQccSampleUpdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateSampleRequest) error {
    r._updateRequest = _updateRequest
    r.Set("update_request", _updateRequest)
    return nil
}

// UpdateRequest Getter
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetUpdateRequest() *UpdateSampleRequest {
    return r._updateRequest
}
