package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代填办理人信息 API请求
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息
*/
type TaobaoAlitripTravelNormalvisaStoreuserRequest struct {
    model.Params
    // 订单id
    _bizOrderId   int64
    // 列表：签证人信息列表
    _normalVisaUserUnitList   []NormalVisaUserUnit
}

// 初始化TaobaoAlitripTravelNormalvisaStoreuserRequest对象
func NewTaobaoAlitripTravelNormalvisaStoreuserRequest() *TaobaoAlitripTravelNormalvisaStoreuserRequest{
    return &TaobaoAlitripTravelNormalvisaStoreuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.storeuser"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaStoreuserRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// NormalVisaUserUnitList Setter
// 列表：签证人信息列表
func (r *TaobaoAlitripTravelNormalvisaStoreuserRequest) SetNormalVisaUserUnitList(_normalVisaUserUnitList []NormalVisaUserUnit) error {
    r._normalVisaUserUnitList = _normalVisaUserUnitList
    r.Set("normal_visa_user_unit_list", _normalVisaUserUnitList)
    return nil
}

// NormalVisaUserUnitList Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserRequest) GetNormalVisaUserUnitList() []NormalVisaUserUnit {
    return r._normalVisaUserUnitList
}
