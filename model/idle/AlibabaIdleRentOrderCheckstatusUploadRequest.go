package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传验收结果 API请求
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

// 初始化AlibabaIdleRentOrderCheckstatusUploadRequest对象
func NewAlibabaIdleRentOrderCheckstatusUploadRequest() *AlibabaIdleRentOrderCheckstatusUploadRequest{
    return &AlibabaIdleRentOrderCheckstatusUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.checkstatus.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetOrderId() int64 {
    return r.orderId
}
// CheckResult Setter
// 校验结果
func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetCheckResult(checkResult *CheckResultDto) error {
    r.checkResult = checkResult
    r.Set("check_result", checkResult)
    return nil
}

// CheckResult Getter
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetCheckResult() *CheckResultDto {
    return r.checkResult
}
