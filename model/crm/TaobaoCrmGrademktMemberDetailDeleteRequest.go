package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-删除商品等级营销明细 APIRequest
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细
*/
type TaobaoCrmGrademktMemberDetailDeleteRequest struct {
    model.Params

    // 扩展字段
    feather   string 

    // 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    parameter   string 

}

func NewTaobaoCrmGrademktMemberDetailDeleteRequest() *TaobaoCrmGrademktMemberDetailDeleteRequest{
    return &TaobaoCrmGrademktMemberDetailDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGrademktMemberDetailDeleteRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.detail.delete"
}

func (r TaobaoCrmGrademktMemberDetailDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGrademktMemberDetailDeleteRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

func (r TaobaoCrmGrademktMemberDetailDeleteRequest) GetFeather() string {
    return r.feather
}

func (r *TaobaoCrmGrademktMemberDetailDeleteRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r TaobaoCrmGrademktMemberDetailDeleteRequest) GetParameter() string {
    return r.parameter
}

