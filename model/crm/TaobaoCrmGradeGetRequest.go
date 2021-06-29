package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询等级规则 API请求
taobao.crm.grade.get

卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
*/
type TaobaoCrmGradeGetRequest struct {
    model.Params
}

// 初始化TaobaoCrmGradeGetRequest对象
func NewTaobaoCrmGradeGetRequest() *TaobaoCrmGradeGetRequest{
    return &TaobaoCrmGradeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGradeGetRequest) GetApiMethodName() string {
    return "taobao.crm.grade.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
