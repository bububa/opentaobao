package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新签证办理进度 API请求
taobao.alitrip.travel.normalvisa.updatepersonstauts

更新签证办理进度
*/
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest struct {
    model.Params
    // 更新信息
    _normalVisaUpdateUnitList   []NormalVisaUpdateUnit
    // 订单号
    _bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest对象
func NewTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest() *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest{
    return &TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.updatepersonstauts"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NormalVisaUpdateUnitList Setter
// 更新信息
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) SetNormalVisaUpdateUnitList(_normalVisaUpdateUnitList []NormalVisaUpdateUnit) error {
    r._normalVisaUpdateUnitList = _normalVisaUpdateUnitList
    r.Set("normal_visa_update_unit_list", _normalVisaUpdateUnitList)
    return nil
}

// NormalVisaUpdateUnitList Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetNormalVisaUpdateUnitList() []NormalVisaUpdateUnit {
    return r._normalVisaUpdateUnitList
}
// BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}