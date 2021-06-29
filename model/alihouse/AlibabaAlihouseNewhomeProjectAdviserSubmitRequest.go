package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘顾问 API请求
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserSubmitRequest struct {
    model.Params
    // 顾问列表
    advisers   []ProjectAdviserDto
}

// 初始化AlibabaAlihouseNewhomeProjectAdviserSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest() *AlibabaAlihouseNewhomeProjectAdviserSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectAdviserSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.adviser.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Advisers Setter
// 顾问列表
func (r *AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) SetAdvisers(advisers []ProjectAdviserDto) error {
    r.advisers = advisers
    r.Set("advisers", advisers)
    return nil
}

// Advisers Getter
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitRequest) GetAdvisers() []ProjectAdviserDto {
    return r.advisers
}
