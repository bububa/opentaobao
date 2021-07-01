package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步供应商校准后的配额-二级物料 API请求
alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload

同步供应商校准后的配额-二级物料
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest struct {
    model.Params
    // 对象
    _currentQuoteRawRequest   *AbstractRequest
}

// 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest{
    return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentQuoteRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) SetCurrentQuoteRawRequest(_currentQuoteRawRequest *AbstractRequest) error {
    r._currentQuoteRawRequest = _currentQuoteRawRequest
    r.Set("current_quote_raw_request", _currentQuoteRawRequest)
    return nil
}

// CurrentQuoteRawRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest) GetCurrentQuoteRawRequest() *AbstractRequest {
    return r._currentQuoteRawRequest
}
