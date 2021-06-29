package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘上下架 APIRequest
alibaba.alihouse.newhome.project.line

上下架楼盘
*/
type AlibabaAlihouseNewhomeProjectLineRequest struct {
    model.Params

    // 外部id
    outerId   string 

    // 0-下架 1-上架
    type   *model.File 

}

func NewAlibabaAlihouseNewhomeProjectLineRequest() *AlibabaAlihouseNewhomeProjectLineRequest{
    return &AlibabaAlihouseNewhomeProjectLineRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectLineRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.line"
}

func (r AlibabaAlihouseNewhomeProjectLineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectLineRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectLineRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaAlihouseNewhomeProjectLineRequest) SetType(type *model.File) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectLineRequest) GetType() *model.File {
    return r.type
}

