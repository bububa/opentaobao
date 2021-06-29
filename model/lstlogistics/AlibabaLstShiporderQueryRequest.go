package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通发货单查询 API请求
alibaba.lst.shiporder.query

通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
*/
type AlibabaLstShiporderQueryRequest struct {
    model.Params
    // 零售通
    source   string
    // 订单
    outOrderId   string
}

// 初始化AlibabaLstShiporderQueryRequest对象
func NewAlibabaLstShiporderQueryRequest() *AlibabaLstShiporderQueryRequest{
    return &AlibabaLstShiporderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.shiporder.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Source Setter
// 零售通
func (r *AlibabaLstShiporderQueryRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaLstShiporderQueryRequest) GetSource() string {
    return r.source
}
// OutOrderId Setter
// 订单
func (r *AlibabaLstShiporderQueryRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaLstShiporderQueryRequest) GetOutOrderId() string {
    return r.outOrderId
}
