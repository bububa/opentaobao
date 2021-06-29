package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘快讯删除 APIRequest
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除
*/
type AlibabaAlihouseNewhomeProjectDynamicDeleteRequest struct {
    model.Params

    // 外部动态ID
    outerDynamicId   string 

}

func NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest() *AlibabaAlihouseNewhomeProjectDynamicDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectDynamicDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.dynamic.delete"
}

func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) SetOuterDynamicId(outerDynamicId string) error {
    r.outerDynamicId = outerDynamicId
    r.Set("outer_dynamic_id", outerDynamicId)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetOuterDynamicId() string {
    return r.outerDynamicId
}

