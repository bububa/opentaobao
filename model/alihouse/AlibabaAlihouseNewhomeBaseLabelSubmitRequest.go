package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标签库 API请求
alibaba.alihouse.newhome.base.label.submit

提交标签库
*/
type AlibabaAlihouseNewhomeBaseLabelSubmitRequest struct {
    model.Params
    // 标签列表
    labels   []BaseLabelDto
}

// 初始化AlibabaAlihouseNewhomeBaseLabelSubmitRequest对象
func NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest() *AlibabaAlihouseNewhomeBaseLabelSubmitRequest{
    return &AlibabaAlihouseNewhomeBaseLabelSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.base.label.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Labels Setter
// 标签列表
func (r *AlibabaAlihouseNewhomeBaseLabelSubmitRequest) SetLabels(labels []BaseLabelDto) error {
    r.labels = labels
    r.Set("labels", labels)
    return nil
}

// Labels Getter
func (r AlibabaAlihouseNewhomeBaseLabelSubmitRequest) GetLabels() []BaseLabelDto {
    return r.labels
}
