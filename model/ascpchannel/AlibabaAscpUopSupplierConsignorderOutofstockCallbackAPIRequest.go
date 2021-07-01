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
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest struct {
    model.Params
    // 缺货回告请求模型
    _consignorderOutofstockCallbackRequest   *Consignorderoutofstockcallbackrequest
}

// 初始化AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest() *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest{
    return &AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.outofstock.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignorderOutofstockCallbackRequest Setter
// 缺货回告请求模型
func (r *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) SetConsignorderOutofstockCallbackRequest(_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest) error {
    r._consignorderOutofstockCallbackRequest = _consignorderOutofstockCallbackRequest
    r.Set("consignorder_outofstock_callback_request", _consignorderOutofstockCallbackRequest)
    return nil
}

// ConsignorderOutofstockCallbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetConsignorderOutofstockCallbackRequest() *Consignorderoutofstockcallbackrequest {
    return r._consignorderOutofstockCallbackRequest
}
