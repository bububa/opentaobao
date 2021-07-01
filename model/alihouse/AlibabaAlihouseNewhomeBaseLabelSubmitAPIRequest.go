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
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest struct {
    model.Params
    // 标签列表
    _labels   []BaseLabelDto
}

// 初始化AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeBaseLabelSubmitRequest() *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest{
    return &AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.base.label.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Labels Setter
// 标签列表
func (r *AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) SetLabels(_labels []BaseLabelDto) error {
    r._labels = _labels
    r.Set("labels", _labels)
    return nil
}

// Labels Getter
func (r AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest) GetLabels() []BaseLabelDto {
    return r._labels
}
