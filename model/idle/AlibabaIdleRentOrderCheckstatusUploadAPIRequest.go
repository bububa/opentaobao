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
type AlibabaIdleRentOrderCheckstatusUploadAPIRequest struct {
    model.Params
    // 订单id
    _orderId   int64
    // 校验结果
    _checkResult   *CheckResultDto
}

// 初始化AlibabaIdleRentOrderCheckstatusUploadAPIRequest对象
func NewAlibabaIdleRentOrderCheckstatusUploadRequest() *AlibabaIdleRentOrderCheckstatusUploadAPIRequest{
    return &AlibabaIdleRentOrderCheckstatusUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderCheckstatusUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.checkstatus.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderCheckstatusUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderCheckstatusUploadAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderCheckstatusUploadAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// CheckResult Setter
// 校验结果
func (r *AlibabaIdleRentOrderCheckstatusUploadAPIRequest) SetCheckResult(_checkResult *CheckResultDto) error {
    r._checkResult = _checkResult
    r.Set("check_result", _checkResult)
    return nil
}

// CheckResult Getter
func (r AlibabaIdleRentOrderCheckstatusUploadAPIRequest) GetCheckResult() *CheckResultDto {
    return r._checkResult
}
