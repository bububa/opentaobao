package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商提交受理订单 APIRequest
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等
*/
type AlibabaTianjiDistributorOrderSubmitRequest struct {
    model.Params

    // 商品编码，如手机串号
    itemSerialNo   string 

    // 淘宝交易订单号
    orderNo   string 

    // 供应商产品编码，如SIM卡号
    productSerialNo   string 

}

func NewAlibabaTianjiDistributorOrderSubmitRequest() *AlibabaTianjiDistributorOrderSubmitRequest{
    return &AlibabaTianjiDistributorOrderSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTianjiDistributorOrderSubmitRequest) GetApiMethodName() string {
    return "alibaba.tianji.distributor.order.submit"
}

func (r AlibabaTianjiDistributorOrderSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetItemSerialNo(itemSerialNo string) error {
    r.itemSerialNo = itemSerialNo
    r.Set("item_serial_no", itemSerialNo)
    return nil
}

func (r AlibabaTianjiDistributorOrderSubmitRequest) GetItemSerialNo() string {
    return r.itemSerialNo
}

func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r AlibabaTianjiDistributorOrderSubmitRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetProductSerialNo(productSerialNo string) error {
    r.productSerialNo = productSerialNo
    r.Set("product_serial_no", productSerialNo)
    return nil
}

func (r AlibabaTianjiDistributorOrderSubmitRequest) GetProductSerialNo() string {
    return r.productSerialNo
}

