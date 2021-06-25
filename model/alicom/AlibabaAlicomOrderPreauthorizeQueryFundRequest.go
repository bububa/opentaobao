package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金流水查询 APIRequest
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFundRequest struct {
    model.Params

    // 入参
    preAuthorizeModel   *PreAuthorizeModel 

}

func NewAlibabaAlicomOrderPreauthorizeQueryFundRequest() *AlibabaAlicomOrderPreauthorizeQueryFundRequest{
    return &AlibabaAlicomOrderPreauthorizeQueryFundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomOrderPreauthorizeQueryFundRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.preauthorize.query.fund"
}

func (r AlibabaAlicomOrderPreauthorizeQueryFundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomOrderPreauthorizeQueryFundRequest) SetPreAuthorizeModel(preAuthorizeModel *PreAuthorizeModel) error {
    r.preAuthorizeModel = preAuthorizeModel
    r.Set("pre_authorize_model", preAuthorizeModel)
    return nil
}

func (r AlibabaAlicomOrderPreauthorizeQueryFundRequest) GetPreAuthorizeModel() *PreAuthorizeModel {
    return r.preAuthorizeModel
}

