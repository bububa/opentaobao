package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家推单 API请求
alibaba.ascp.industry.uop.supplier.consignoder

商家推单
*/
type AlibabaAscpIndustryUopSupplierConsignoderRequest struct {
    model.Params
    // 发货主单信息
    _erpNormalConsignOrderRequest   *Erpnormalconsignorderrequest
}

// 初始化AlibabaAscpIndustryUopSupplierConsignoderRequest对象
func NewAlibabaAscpIndustryUopSupplierConsignoderRequest() *AlibabaAscpIndustryUopSupplierConsignoderRequest{
    return &AlibabaAscpIndustryUopSupplierConsignoderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.uop.supplier.consignoder"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErpNormalConsignOrderRequest Setter
// 发货主单信息
func (r *AlibabaAscpIndustryUopSupplierConsignoderRequest) SetErpNormalConsignOrderRequest(_erpNormalConsignOrderRequest *Erpnormalconsignorderrequest) error {
    r._erpNormalConsignOrderRequest = _erpNormalConsignOrderRequest
    r.Set("erp_normal_consign_order_request", _erpNormalConsignOrderRequest)
    return nil
}

// ErpNormalConsignOrderRequest Getter
func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetErpNormalConsignOrderRequest() *Erpnormalconsignorderrequest {
    return r._erpNormalConsignOrderRequest
}
