package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitJudgementGetAPIRequest
获取裁判登记信息 API请求
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息 */
type AlibabaLegalSuitJudgementGetAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
}

// NewAlibabaLegalSuitJudgementGetRequest 初始化AlibabaLegalSuitJudgementGetAPIRequest对象
func NewAlibabaLegalSuitJudgementGetRequest() *AlibabaLegalSuitJudgementGetAPIRequest {
	return &AlibabaLegalSuitJudgementGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SuitId Setter
// 案件id
func (r *AlibabaLegalSuitJudgementGetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// Get SuitId Getter
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetSuitId() int64 {
	return r._suitId
}
