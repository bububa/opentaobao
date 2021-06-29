package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员等级营销-等级营销活动查询 API请求
taobao.crm.grademkt.member.detail.query

商家通过该接口查询等级营销活动
*/
type TaobaoCrmGrademktMemberDetailQueryRequest struct {
    model.Params
    // 扩展字段
    feather   string
    // 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
    parameter   string
}

// 初始化TaobaoCrmGrademktMemberDetailQueryRequest对象
func NewTaobaoCrmGrademktMemberDetailQueryRequest() *TaobaoCrmGrademktMemberDetailQueryRequest{
    return &TaobaoCrmGrademktMemberDetailQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailQueryRequest) GetApiMethodName() string {
    return "taobao.crm.grademkt.member.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailQueryRequest) SetFeather(feather string) error {
    r.feather = feather
    r.Set("feather", feather)
    return nil
}

// Feather Getter
func (r TaobaoCrmGrademktMemberDetailQueryRequest) GetFeather() string {
    return r.feather
}
// Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailQueryRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

// Parameter Getter
func (r TaobaoCrmGrademktMemberDetailQueryRequest) GetParameter() string {
    return r.parameter
}
