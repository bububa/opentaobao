package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单纬度的仓缺货回告服务 API请求
alibaba.ascp.uop.supplier.consignorder.outofstock.callback

商家仓履约单纬度的仓缺货回告接口
*/
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest struct {
    model.Params
    // 缺货回告请求模型
    _consignorderOutofstockCallbackRequest   *Consignorderoutofstockcallbackrequest
}

// 初始化AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest对象
func NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest() *AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest{
    return &AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.outofstock.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignorderOutofstockCallbackRequest Setter
// 缺货回告请求模型
func (r *AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) SetConsignorderOutofstockCallbackRequest(_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest) error {
    r._consignorderOutofstockCallbackRequest = _consignorderOutofstockCallbackRequest
    r.Set("consignorder_outofstock_callback_request", _consignorderOutofstockCallbackRequest)
    return nil
}

// ConsignorderOutofstockCallbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetConsignorderOutofstockCallbackRequest() *Consignorderoutofstockcallbackrequest {
    return r._consignorderOutofstockCallbackRequest
}
