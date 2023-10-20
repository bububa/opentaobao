package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitjudgementgetAPIRequest 获取裁判登记信息 API请求
// alibaba.legal.suit.judgement.get
//
// 供ISV供应商获取集团法务系统的裁判登记信息
type AlibabalegalsuitjudgementgetAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
}

// NewAlibabalegalsuitjudgementgetRequest 初始化AlibabalegalsuitjudgementgetAPIRequest对象
func NewAlibabalegalsuitjudgementgetRequest() *AlibabalegalsuitjudgementgetAPIRequest {
	return &AlibabalegalsuitjudgementgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitjudgementgetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitjudgementgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitjudgementgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件id
func (r *AlibabalegalsuitjudgementgetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabalegalsuitjudgementgetAPIRequest) GetSuitId() int64 {
	return r._suitId
}
