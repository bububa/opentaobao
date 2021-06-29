package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘预售证 APIRequest
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest struct {
    model.Params

    // 外部顾问ID
    outerPermitId   string 

}

func NewAlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.delete"
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) SetOuterPermitId(outerPermitId string) error {
    r.outerPermitId = outerPermitId
    r.Set("outer_permit_id", outerPermitId)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetOuterPermitId() string {
    return r.outerPermitId
}

