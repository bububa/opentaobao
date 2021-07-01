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
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest struct {
    model.Params
    // 外部动态ID
    _outerDynamicId   string
}

// 初始化AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest() *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest{
    return &AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.dynamic.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterDynamicId Setter
// 外部动态ID
func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) SetOuterDynamicId(_outerDynamicId string) error {
    r._outerDynamicId = _outerDynamicId
    r.Set("outer_dynamic_id", _outerDynamicId)
    return nil
}

// OuterDynamicId Getter
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetOuterDynamicId() string {
    return r._outerDynamicId
}
