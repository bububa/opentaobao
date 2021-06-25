package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家设置等级规则 APIRequest
taobao.crm.grade.set

设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
*/
type TaobaoCrmGradeSetRequest struct {
    model.Params

    // 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br><br/>amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。
    amount   []Number 

    // 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br><br/>count参数的填写规范：<br/>升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。
    count   []Number 

    // 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP
    grade   []Number 

    // 会员级别折扣率。会员等级越高，折扣必须越低。<br/>950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。
    discount   []Number 

    // 是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额
    hierarchy   []Boolean 

}

func NewTaobaoCrmGradeSetRequest() *TaobaoCrmGradeSetRequest{
    return &TaobaoCrmGradeSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGradeSetRequest) GetApiMethodName() string {
    return "taobao.crm.grade.set"
}

func (r TaobaoCrmGradeSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGradeSetRequest) SetAmount(amount []Number) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r TaobaoCrmGradeSetRequest) GetAmount() []Number {
    return r.amount
}

func (r *TaobaoCrmGradeSetRequest) SetCount(count []Number) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r TaobaoCrmGradeSetRequest) GetCount() []Number {
    return r.count
}

func (r *TaobaoCrmGradeSetRequest) SetGrade(grade []Number) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

func (r TaobaoCrmGradeSetRequest) GetGrade() []Number {
    return r.grade
}

func (r *TaobaoCrmGradeSetRequest) SetDiscount(discount []Number) error {
    r.discount = discount
    r.Set("discount", discount)
    return nil
}

func (r TaobaoCrmGradeSetRequest) GetDiscount() []Number {
    return r.discount
}

func (r *TaobaoCrmGradeSetRequest) SetHierarchy(hierarchy []Boolean) error {
    r.hierarchy = hierarchy
    r.Set("hierarchy", hierarchy)
    return nil
}

func (r TaobaoCrmGradeSetRequest) GetHierarchy() []Boolean {
    return r.hierarchy
}

