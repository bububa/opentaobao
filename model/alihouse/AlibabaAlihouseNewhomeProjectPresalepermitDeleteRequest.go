package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘预售证 API请求
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest struct {
    model.Params
    // 外部顾问ID
    outerPermitId   string
}

// 初始化AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterPermitId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) SetOuterPermitId(outerPermitId string) error {
    r.outerPermitId = outerPermitId
    r.Set("outer_permit_id", outerPermitId)
    return nil
}

// OuterPermitId Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest) GetOuterPermitId() string {
    return r.outerPermitId
}
