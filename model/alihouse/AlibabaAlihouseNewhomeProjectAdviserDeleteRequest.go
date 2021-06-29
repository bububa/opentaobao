package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘顾问 APIRequest
alibaba.alihouse.newhome.project.adviser.delete

删除楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserDeleteRequest struct {
    model.Params

    // 外部顾问ID
    outerConsultantId   string 

}

func NewAlibabaAlihouseNewhomeProjectAdviserDeleteRequest() *AlibabaAlihouseNewhomeProjectAdviserDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectAdviserDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.adviser.delete"
}

func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) SetOuterConsultantId(outerConsultantId string) error {
    r.outerConsultantId = outerConsultantId
    r.Set("outer_consultant_id", outerConsultantId)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetOuterConsultantId() string {
    return r.outerConsultantId
}

