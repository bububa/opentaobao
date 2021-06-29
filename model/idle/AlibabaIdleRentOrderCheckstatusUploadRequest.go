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
    _orderId   int64
    // 校验结果
    _checkResult   *CheckResultDTO
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
func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetOrderId() int64 {
    return r._orderId
}
// CheckResult Setter
// 校验结果
func (r *AlibabaIdleRentOrderCheckstatusUploadRequest) SetCheckResult(_checkResult *CheckResultDTO) error {
    r._checkResult = _checkResult
    r.Set("check_result", _checkResult)
    return nil
}

// CheckResult Getter
func (r AlibabaIdleRentOrderCheckstatusUploadRequest) GetCheckResult() *CheckResultDTO {
    return r._checkResult
}
