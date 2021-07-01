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
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest struct {
    model.Params
    // 外部顾问ID
    _outerPermitId   string
}

// 初始化AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest{
    return &AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.presalepermit.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterPermitId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) SetOuterPermitId(_outerPermitId string) error {
    r._outerPermitId = _outerPermitId
    r.Set("outer_permit_id", _outerPermitId)
    return nil
}

// OuterPermitId Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetOuterPermitId() string {
    return r._outerPermitId
}
