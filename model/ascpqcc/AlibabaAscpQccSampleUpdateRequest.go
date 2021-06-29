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
type AlibabaAscpQccSampleUpdateRequest struct {
    model.Params
    // 更新请求参数
    updateRequest   *UpdateSampleRequest
}

// 初始化AlibabaAscpQccSampleUpdateRequest对象
func NewAlibabaAscpQccSampleUpdateRequest() *AlibabaAscpQccSampleUpdateRequest{
    return &AlibabaAscpQccSampleUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.qcc.sample.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRequest Setter
// 更新请求参数
func (r *AlibabaAscpQccSampleUpdateRequest) SetUpdateRequest(updateRequest *UpdateSampleRequest) error {
    r.updateRequest = updateRequest
    r.Set("update_request", updateRequest)
    return nil
}

// UpdateRequest Getter
func (r AlibabaAscpQccSampleUpdateRequest) GetUpdateRequest() *UpdateSampleRequest {
    return r.updateRequest
}
