package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商提交受理订单 API请求
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等
*/
type AlibabaTianjiDistributorOrderSubmitRequest struct {
    model.Params
    // 商品编码，如手机串号
    _itemSerialNo   string
    // 淘宝交易订单号
    _orderNo   string
    // 供应商产品编码，如SIM卡号
    _productSerialNo   string
}

// 初始化AlibabaTianjiDistributorOrderSubmitRequest对象
func NewAlibabaTianjiDistributorOrderSubmitRequest() *AlibabaTianjiDistributorOrderSubmitRequest{
    return &AlibabaTianjiDistributorOrderSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTianjiDistributorOrderSubmitRequest) GetApiMethodName() string {
    return "alibaba.tianji.distributor.order.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTianjiDistributorOrderSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemSerialNo Setter
// 商品编码，如手机串号
func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetItemSerialNo(_itemSerialNo string) error {
    r._itemSerialNo = _itemSerialNo
    r.Set("item_serial_no", _itemSerialNo)
    return nil
}

// ItemSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitRequest) GetItemSerialNo() string {
    return r._itemSerialNo
}
// OrderNo Setter
// 淘宝交易订单号
func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaTianjiDistributorOrderSubmitRequest) GetOrderNo() string {
    return r._orderNo
}
// ProductSerialNo Setter
// 供应商产品编码，如SIM卡号
func (r *AlibabaTianjiDistributorOrderSubmitRequest) SetProductSerialNo(_productSerialNo string) error {
    r._productSerialNo = _productSerialNo
    r.Set("product_serial_no", _productSerialNo)
    return nil
}

// ProductSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitRequest) GetProductSerialNo() string {
    return r._productSerialNo
}
