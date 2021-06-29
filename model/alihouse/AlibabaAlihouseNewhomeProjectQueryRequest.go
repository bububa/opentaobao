package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询楼盘相关信息 APIRequest
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息
*/
type AlibabaAlihouseNewhomeProjectQueryRequest struct {
    model.Params

    // 外部楼盘/小区id
    outerId   string 

}

func NewAlibabaAlihouseNewhomeProjectQueryRequest() *AlibabaAlihouseNewhomeProjectQueryRequest{
    return &AlibabaAlihouseNewhomeProjectQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.query"
}

func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectQueryRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetOuterId() string {
    return r.outerId
}

