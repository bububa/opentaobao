package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向销退入库单入库结果回告 APIRequest
alibaba.ascp.uop.supplier.reverseorder.instorage.feedback

ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
*/
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest struct {
    model.Params

    // 销退单入库结果请求
    instorageFeedbackRequest   *Instoragefeedbackrequest 

}

func NewAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest() *AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest{
    return &AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.reverseorder.instorage.feedback"
}

func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) SetInstorageFeedbackRequest(instorageFeedbackRequest *Instoragefeedbackrequest) error {
    r.instorageFeedbackRequest = instorageFeedbackRequest
    r.Set("instorage_feedback_request", instorageFeedbackRequest)
    return nil
}

func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest) GetInstorageFeedbackRequest() *Instoragefeedbackrequest {
    return r.instorageFeedbackRequest
}

