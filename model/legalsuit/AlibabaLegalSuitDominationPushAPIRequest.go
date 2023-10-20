package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitdominationpushAPIRequest 更新或者保存管辖信息 API请求
// alibaba.legal.suit.domination.push
//
// ISV推送管辖信息到诉讼平台
type AlibabalegalsuitdominationpushAPIRequest struct {
	model.Params
	// 管辖信息
	_dominationDissentModel *DominationDissentModel
}

// NewAlibabalegalsuitdominationpushRequest 初始化AlibabalegalsuitdominationpushAPIRequest对象
func NewAlibabalegalsuitdominationpushRequest() *AlibabalegalsuitdominationpushAPIRequest {
	return &AlibabalegalsuitdominationpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitdominationpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.domination.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitdominationpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitdominationpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDominationDissentModel is DominationDissentModel Setter
// 管辖信息
func (r *AlibabalegalsuitdominationpushAPIRequest) SetDominationDissentModel(_dominationDissentModel *DominationDissentModel) error {
	r._dominationDissentModel = _dominationDissentModel
	r.Set("domination_dissent_model", _dominationDissentModel)
	return nil
}

// GetDominationDissentModel DominationDissentModel Getter
func (r AlibabalegalsuitdominationpushAPIRequest) GetDominationDissentModel() *DominationDissentModel {
	return r._dominationDissentModel
}
