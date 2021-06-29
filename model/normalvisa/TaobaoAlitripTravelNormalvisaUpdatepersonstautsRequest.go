package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新签证办理进度 APIRequest
taobao.alitrip.travel.normalvisa.updatepersonstauts

更新签证办理进度
*/
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest struct {
    model.Params

    // 更新信息
    normalVisaUpdateUnitList   []NormalVisaUpdateUnit 

    // 订单号
    bizOrderId   int64 

}

func NewTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest() *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest{
    return &TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.updatepersonstauts"
}

func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) SetNormalVisaUpdateUnitList(normalVisaUpdateUnitList []NormalVisaUpdateUnit) error {
    r.normalVisaUpdateUnitList = normalVisaUpdateUnitList
    r.Set("normal_visa_update_unit_list", normalVisaUpdateUnitList)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetNormalVisaUpdateUnitList() []NormalVisaUpdateUnit {
    return r.normalVisaUpdateUnitList
}

func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

