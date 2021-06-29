package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者保存管辖信息 API请求
alibaba.legal.suit.domination.push

ISV推送管辖信息到诉讼平台
*/
type AlibabaLegalSuitDominationPushRequest struct {
    model.Params
    // 管辖信息
    _dominationDissentModel   *DominationDissentModel
}

// 初始化AlibabaLegalSuitDominationPushRequest对象
func NewAlibabaLegalSuitDominationPushRequest() *AlibabaLegalSuitDominationPushRequest{
    return &AlibabaLegalSuitDominationPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitDominationPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.domination.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitDominationPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DominationDissentModel Setter
// 管辖信息
func (r *AlibabaLegalSuitDominationPushRequest) SetDominationDissentModel(_dominationDissentModel *DominationDissentModel) error {
    r._dominationDissentModel = _dominationDissentModel
    r.Set("domination_dissent_model", _dominationDissentModel)
    return nil
}

// DominationDissentModel Getter
func (r AlibabaLegalSuitDominationPushRequest) GetDominationDissentModel() *DominationDissentModel {
    return r._dominationDissentModel
}
