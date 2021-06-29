package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-创建商品等级营销明细 API请求
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

// 初始化TaobaoCrmGrademktMemberDetailCreateRequest对象
func NewTaobaoCrmGrademktMemberDetailCreateRequest() *TaobaoCrmGrademktMemberDetailCreateRequest{
    return &TaobaoCrmGrademktMemberDetailCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.detail.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailCreateRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetFeather() string {
    return r.feather
}
// Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailCreateRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberDetailCreateRequest) GetParameter() string {
    return r.parameter
}
