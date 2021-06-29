package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传验收结果 APIRequest
alibaba.idle.rent.order.checkstatus.upload

上传验收结果
*/
type AlibabaIdleRentOrderCheckstatusUploadRequest struct {
    model.Params

    // 订单id
    orderId   int64 

    // 校验结果
    checkResult   *CheckResultDto 

}

func NewAlibabaIdleRentOrderCheckstatusUploadRequest() *AlibabaIdleRentOrderCheckstatusUploadRequest{
    return &AlibabaIdleRentOrderCheckstatusUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.checkstatus.upload"
}

func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetCheckResult(checkResult *CheckResultDto) error {
    r.checkResult = checkResult
    r.Set("check_result", checkResult)
    return nil
}

func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetCheckResult() *CheckResultDto {
    return r.checkResult
}

