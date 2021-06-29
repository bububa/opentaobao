package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘顾问 API请求
alibaba.alihouse.newhome.project.adviser.delete

删除楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserDeleteRequest struct {
    model.Params
    // 外部顾问ID
    _outerConsultantId   string
}

// 初始化AlibabaAlihouseNewhomeProjectAdviserDeleteRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserDeleteRequest() *AlibabaAlihouseNewhomeProjectAdviserDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectAdviserDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.adviser.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterConsultantId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) SetOuterConsultantId(_outerConsultantId string) error {
    r._outerConsultantId = _outerConsultantId
    r.Set("outer_consultant_id", _outerConsultantId)
    return nil
}

// OuterConsultantId Getter
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteRequest) GetOuterConsultantId() string {
    return r._outerConsultantId
}
