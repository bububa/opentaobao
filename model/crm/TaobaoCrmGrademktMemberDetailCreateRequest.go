package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-创建商品等级营销明细 APIRequest
taobao.crm.grademkt.member.detail.create

创建商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailCreateRequest struct {
    model.Params

    // 扩展字段
    feather   string 

    // 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    parameter   string 

}

func NewTaobaoCrmGrademktMemberDetailCreateRequest() *TaobaoCrmGrademktMemberDetailCreateRequest{
    return &TaobaoCrmGrademktMemberDetailCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.detail.create"
}

func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGrademktMemberDetailCreateRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetFeather() string {
    return r.feather
}

func (r *TaobaoCrmGrademktMemberDetailCreateRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetParameter() string {
    return r.parameter
}

