package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家仓wms取消发货反馈回告服务 API请求
alibaba.ascp.uop.supplier.consignorder.cancel.feedback

履约单纬度通知商家仓wms取消发货结果反馈回告服务
*/
type AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest struct {
    model.Params
    // 取消发货反馈回告请求
    consignorderCancelFeedbackRequest   *Consignordercancelfeedbackrequest
}

// 初始化AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest对象
func NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest() *AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest{
    return &AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.cancel.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignorderCancelFeedbackRequest Setter
// 取消发货反馈回告请求
func (r *AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) SetConsignorderCancelFeedbackRequest(consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest) error {
    r.consignorderCancelFeedbackRequest = consignorderCancelFeedbackRequest
    r.Set("consignorder_cancel_feedback_request", consignorderCancelFeedbackRequest)
    return nil
}

// ConsignorderCancelFeedbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest) GetConsignorderCancelFeedbackRequest() *Consignordercancelfeedbackrequest {
    return r.consignorderCancelFeedbackRequest
}
