package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务办理结果 API请求
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知
*/
type AlibabaAlicomOrderPreauthorizeCreateRequest struct {
    model.Params
    // 入参
    _preAuthorizeModel   *PreAuthorizeModel
}

// 初始化AlibabaAlicomOrderPreauthorizeCreateRequest对象
func NewAlibabaAlicomOrderPreauthorizeCreateRequest() *AlibabaAlicomOrderPreauthorizeCreateRequest{
    return &AlibabaAlicomOrderPreauthorizeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreAuthorizeModel Setter
// 入参
func (r *AlibabaAlicomOrderPreauthorizeCreateRequest) SetPreAuthorizeModel(_preAuthorizeModel *PreAuthorizeModel) error {
    r._preAuthorizeModel = _preAuthorizeModel
    r.Set("pre_authorize_model", _preAuthorizeModel)
    return nil
}

// PreAuthorizeModel Getter
func (r AlibabaAlicomOrderPreauthorizeCreateRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r._preAuthorizeModel
}
