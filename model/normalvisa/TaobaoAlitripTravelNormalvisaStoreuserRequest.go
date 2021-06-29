package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代填办理人信息 APIRequest
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息
*/
type TaobaoAlitripTravelNormalvisaStoreuserRequest struct {
    model.Params

    // 订单id
    bizOrderId   int64 

    // 列表：签证人信息列表
    normalVisaUserUnitList   []NormalVisaUserUnit 

}

func NewTaobaoAlitripTravelNormalvisaStoreuserRequest() *TaobaoAlitripTravelNormalvisaStoreuserRequest{
    return &TaobaoAlitripTravelNormalvisaStoreuserRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.storeuser"
}

func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaStoreuserRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *TaobaoAlitripTravelNormalvisaStoreuserRequest) SetNormalVisaUserUnitList(normalVisaUserUnitList []NormalVisaUserUnit) error {
    r.normalVisaUserUnitList = normalVisaUserUnitList
    r.Set("normal_visa_user_unit_list", normalVisaUserUnitList)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetNormalVisaUserUnitList() []NormalVisaUserUnit {
    return r.normalVisaUserUnitList
}

