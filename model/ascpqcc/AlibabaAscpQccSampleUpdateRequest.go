package ascpqcc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品控中心更新样品信息 APIRequest
alibaba.ascp.qcc.sample.update

品控中心更新样品信息
*/
type AlibabaAscpQccSampleUpdateRequest struct {
    model.Params

    // 更新请求参数
    updateRequest   *UpdateSampleRequest 

}

func NewAlibabaAscpQccSampleUpdateRequest() *AlibabaAscpQccSampleUpdateRequest{
    return &AlibabaAscpQccSampleUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpQccSampleUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.qcc.sample.update"
}

func (r AlibabaAscpQccSampleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpQccSampleUpdateRequest) SetUpdateRequest(updateRequest *UpdateSampleRequest) error {
    r.updateRequest = updateRequest
    r.Set("update_request", updateRequest)
    return nil
}

func (r AlibabaAscpQccSampleUpdateRequest) GetUpdateRequest() *UpdateSampleRequest {
    return r.updateRequest
}

