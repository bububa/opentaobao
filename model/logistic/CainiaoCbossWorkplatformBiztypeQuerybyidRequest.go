package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据业务类型id查询业务类型详细信息 API请求
cainiao.cboss.workplatform.biztype.querybyid

菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQuerybyidRequest struct {
    model.Params
    // 业务类型id
    _bizTypeId   string
}

// 初始化CainiaoCbossWorkplatformBiztypeQuerybyidRequest对象
func NewCainiaoCbossWorkplatformBiztypeQuerybyidRequest() *CainiaoCbossWorkplatformBiztypeQuerybyidRequest{
    return &CainiaoCbossWorkplatformBiztypeQuerybyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.biztype.querybyid"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizTypeId Setter
// 业务类型id
func (r *CainiaoCbossWorkplatformBiztypeQuerybyidRequest) SetBizTypeId(_bizTypeId string) error {
    r._bizTypeId = _bizTypeId
    r.Set("biz_type_id", _bizTypeId)
    return nil
}

// BizTypeId Getter
func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetBizTypeId() string {
    return r._bizTypeId
}
