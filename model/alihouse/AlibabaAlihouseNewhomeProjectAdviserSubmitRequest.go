package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘顾问 APIRequest
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserSubmitRequest struct {
    model.Params

    // 顾问列表
    advisers   []ProjectAdviserDto 

}

func NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest() *AlibabaAlihouseNewhomeProjectAdviserSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectAdviserSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.adviser.submit"
}

func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) SetAdvisers(advisers []ProjectAdviserDto) error {
    r.advisers = advisers
    r.Set("advisers", advisers)
    return nil
}

func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetAdvisers() []ProjectAdviserDto {
    return r.advisers
}

