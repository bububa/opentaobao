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
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest struct {
    model.Params
    // 对象
    _currentQuoteRawRequest   *AbstractRequest
}

// 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest{
    return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentQuoteRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) SetCurrentQuoteRawRequest(_currentQuoteRawRequest *AbstractRequest) error {
    r._currentQuoteRawRequest = _currentQuoteRawRequest
    r.Set("current_quote_raw_request", _currentQuoteRawRequest)
    return nil
}

// CurrentQuoteRawRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetCurrentQuoteRawRequest() *AbstractRequest {
    return r._currentQuoteRawRequest
}
