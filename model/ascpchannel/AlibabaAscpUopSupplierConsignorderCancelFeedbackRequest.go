package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家仓wms取消发货反馈回告服务 APIRequest
alibaba.ascp.uop.supplier.consignorder.cancel.feedback

履约单纬度通知商家仓wms取消发货结果反馈回告服务
*/
type AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest struct {
    model.Params

    // 取消发货反馈回告请求
    consignorderCancelFeedbackRequest   *Consignordercancelfeedbackrequest 

}

func NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest() *AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest{
    return &AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.cancel.feedback"
}

func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) SetConsignorderCancelFeedbackRequest(consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest) error {
    r.consignorderCancelFeedbackRequest = consignorderCancelFeedbackRequest
    r.Set("consignorder_cancel_feedback_request", consignorderCancelFeedbackRequest)
    return nil
}

func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetConsignorderCancelFeedbackRequest() *Consignordercancelfeedbackrequest {
    return r.consignorderCancelFeedbackRequest
}

