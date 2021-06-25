package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收商同步前置补贴红包的代扣成功事件 APIRequest
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件
*/
type TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest struct {
    model.Params

    // 变化的金额
    deductAmount   int64 

    // 旧机单id
    oldOrderId   int64 

}

func NewTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest() *TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest{
    return &TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) GetApiMethodName() string {
    return "taobao.recycle.ofnpreredpacket.tpdeductsuccess"
}

func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) SetDeductAmount(deductAmount int64) error {
    r.deductAmount = deductAmount
    r.Set("deduct_amount", deductAmount)
    return nil
}

func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) GetDeductAmount() int64 {
    return r.deductAmount
}

func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) SetOldOrderId(oldOrderId int64) error {
    r.oldOrderId = oldOrderId
    r.Set("old_order_id", oldOrderId)
    return nil
}

func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessRequest) GetOldOrderId() int64 {
    return r.oldOrderId
}

