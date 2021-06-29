package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通发货单创建 API请求
alibaba.lst.shiporder.create

通过该接口可以创建零售通运保保发货单，并处理相关业务流程。
*/
type AlibabaLstShiporderCreateRequest struct {
    model.Params
    // 创建发货单入参
    shipOrder   *LstThirdPartMainShipOrderCreateDto
}

// 初始化AlibabaLstShiporderCreateRequest对象
func NewAlibabaLstShiporderCreateRequest() *AlibabaLstShiporderCreateRequest{
    return &AlibabaLstShiporderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderCreateRequest) GetApiMethodName() string {
    return "alibaba.lst.shiporder.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShipOrder Setter
// 创建发货单入参
func (r *AlibabaLstShiporderCreateRequest) SetShipOrder(shipOrder *LstThirdPartMainShipOrderCreateDto) error {
    r.shipOrder = shipOrder
    r.Set("ship_order", shipOrder)
    return nil
}

// ShipOrder Getter
func (r AlibabaLstShiporderCreateRequest) GetShipOrder() *LstThirdPartMainShipOrderCreateDto {
    return r.shipOrder
}
