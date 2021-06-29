package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解决方案服务资源列表 API请求
cainiao.global.solution.service.resource.query

返回直接解决方案的指定物流服务的可用资源列表
*/
type CainiaoGlobalSolutionServiceResourceQueryRequest struct {
    model.Params
    // 多语言信息
    _locale   string
    // 商家信息
    _sellerParam   *SellerParam
    // 查询参数
    _solutionServiceResParam   *QuerySolutionServiceResParam
    // 发件信息
    _senderParam   *OpenSenderParam
}

// 初始化CainiaoGlobalSolutionServiceResourceQueryRequest对象
func NewCainiaoGlobalSolutionServiceResourceQueryRequest() *CainiaoGlobalSolutionServiceResourceQueryRequest{
    return &CainiaoGlobalSolutionServiceResourceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetApiMethodName() string {
    return "cainiao.global.solution.service.resource.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Locale Setter
// 多语言信息
func (r *CainiaoGlobalSolutionServiceResourceQueryRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetLocale() string {
    return r._locale
}
// SellerParam Setter
// 商家信息
func (r *CainiaoGlobalSolutionServiceResourceQueryRequest) SetSellerParam(_sellerParam *SellerParam) error {
    r._sellerParam = _sellerParam
    r.Set("seller_param", _sellerParam)
    return nil
}

// SellerParam Getter
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetSellerParam() *SellerParam {
    return r._sellerParam
}
// SolutionServiceResParam Setter
// 查询参数
func (r *CainiaoGlobalSolutionServiceResourceQueryRequest) SetSolutionServiceResParam(_solutionServiceResParam *QuerySolutionServiceResParam) error {
    r._solutionServiceResParam = _solutionServiceResParam
    r.Set("solution_service_res_param", _solutionServiceResParam)
    return nil
}

// SolutionServiceResParam Getter
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetSolutionServiceResParam() *QuerySolutionServiceResParam {
    return r._solutionServiceResParam
}
// SenderParam Setter
// 发件信息
func (r *CainiaoGlobalSolutionServiceResourceQueryRequest) SetSenderParam(_senderParam *OpenSenderParam) error {
    r._senderParam = _senderParam
    r.Set("sender_param", _senderParam)
    return nil
}

// SenderParam Getter
func (r CainiaoGlobalSolutionServiceResourceQueryRequest) GetSenderParam() *OpenSenderParam {
    return r._senderParam
}
