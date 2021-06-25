package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据业务类型id查询业务类型详细信息 APIRequest
cainiao.cboss.workplatform.biztype.querybyid

菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQuerybyidRequest struct {
    model.Params

    // 业务类型id
    bizTypeId   string 

}

func NewCainiaoCbossWorkplatformBiztypeQuerybyidRequest() *CainiaoCbossWorkplatformBiztypeQuerybyidRequest{
    return &CainiaoCbossWorkplatformBiztypeQuerybyidRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.biztype.querybyid"
}

func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformBiztypeQuerybyidRequest) SetBizTypeId(bizTypeId string) error {
    r.bizTypeId = bizTypeId
    r.Set("biz_type_id", bizTypeId)
    return nil
}

func (r CainiaoCbossWorkplatformBiztypeQuerybyidRequest) GetBizTypeId() string {
    return r.bizTypeId
}

