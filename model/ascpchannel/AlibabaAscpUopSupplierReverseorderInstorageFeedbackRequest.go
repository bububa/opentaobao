package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向销退入库单入库结果回告 API请求
alibaba.ascp.uop.supplier.reverseorder.instorage.feedback

ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
*/
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest struct {
    model.Params
    // 销退单入库结果请求
    _instorageFeedbackRequest   *Instoragefeedbackrequest
}

// 初始化AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest对象
func NewAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest() *AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest{
    return &AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.instorage.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InstorageFeedbackRequest Setter
// 销退单入库结果请求
func (r *AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) SetInstorageFeedbackRequest(_instorageFeedbackRequest *Instoragefeedbackrequest) error {
    r._instorageFeedbackRequest = _instorageFeedbackRequest
    r.Set("instorage_feedback_request", _instorageFeedbackRequest)
    return nil
}

// InstorageFeedbackRequest Getter
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetInstorageFeedbackRequest() *Instoragefeedbackrequest {
    return r._instorageFeedbackRequest
}
