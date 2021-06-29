package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标签库 APIRequest
alibaba.alihouse.newhome.base.label.submit

提交标签库
*/
type AlibabaAlihouseNewhomeBaseLabelSubmitRequest struct {
    model.Params

    // 标签列表
    labels   []BaseLabelDto 

}

func NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest() *AlibabaAlihouseNewhomeBaseLabelSubmitRequest{
    return &AlibabaAlihouseNewhomeBaseLabelSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.base.label.submit"
}

func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeBaseLabelSubmitRequest) SetLabels(labels []BaseLabelDto) error {
    r.labels = labels
    r.Set("labels", labels)
    return nil
}

func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetLabels() []BaseLabelDto {
    return r.labels
}

