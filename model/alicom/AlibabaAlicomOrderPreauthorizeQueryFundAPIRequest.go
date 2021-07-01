package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金流水查询 API请求
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest struct {
    model.Params
    // 入参
    _preAuthorizeModel   *PreAuthorizeModel
}

// 初始化AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest对象
func NewAlibabaAlicomOrderPreauthorizeQueryFundRequest() *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest{
    return &AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.query.fund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
    r._preAuthorizeModel = _preAuthorizeModel
    r.Set("pre_authorize_model", _preAuthorizeModel)
    return nil
}

// PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r._preAuthorizeModel
}
