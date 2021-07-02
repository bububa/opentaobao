package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitDominationPushAPIRequest 更新或者保存管辖信息 API请求
// alibaba.legal.suit.domination.push
//
// ISV推送管辖信息到诉讼平台
type AlibabaLegalSuitDominationPushAPIRequest struct {
	model.Params
	// 管辖信息
	_dominationDissentModel *DominationDissentModel
}

// NewAlibabaLegalSuitDominationPushRequest 初始化AlibabaLegalSuitDominationPushAPIRequest对象
func NewAlibabaLegalSuitDominationPushRequest() *AlibabaLegalSuitDominationPushAPIRequest {
	return &AlibabaLegalSuitDominationPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitDominationPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.domination.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitDominationPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DominationDissentModel Setter
// 管辖信息
func (r *AlibabaLegalSuitDominationPushAPIRequest) SetDominationDissentModel(_dominationDissentModel *DominationDissentModel) error {
	r._dominationDissentModel = _dominationDissentModel
	r.Set("domination_dissent_model", _dominationDissentModel)
	return nil
}

// Get DominationDissentModel Getter
func (r AlibabaLegalSuitDominationPushAPIRequest) GetDominationDissentModel() *DominationDissentModel {
	return r._dominationDissentModel
}
