package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘快讯删除 API请求
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除
*/
type AlibabaAlihouseNewhomeProjectDynamicDeleteRequest struct {
    model.Params
    // 外部动态ID
    outerDynamicId   string
}

// 初始化AlibabaAlihouseNewhomeProjectDynamicDeleteRequest对象
func NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest() *AlibabaAlihouseNewhomeProjectDynamicDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectDynamicDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.dynamic.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterDynamicId Setter
// 外部动态ID
func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) SetOuterDynamicId(outerDynamicId string) error {
    r.outerDynamicId = outerDynamicId
    r.Set("outer_dynamic_id", outerDynamicId)
    return nil
}

// OuterDynamicId Getter
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteRequest) GetOuterDynamicId() string {
    return r.outerDynamicId
}
