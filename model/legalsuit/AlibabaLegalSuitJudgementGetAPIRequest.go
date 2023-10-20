package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitJudgementGetAPIRequest 获取裁判登记信息 API请求
// alibaba.legal.suit.judgement.get
//
// 供ISV供应商获取集团法务系统的裁判登记信息
type AlibabaLegalSuitJudgementGetAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
}

// NewAlibabaLegalSuitJudgementGetRequest 初始化AlibabaLegalSuitJudgementGetAPIRequest对象
func NewAlibabaLegalSuitJudgementGetRequest() *AlibabaLegalSuitJudgementGetAPIRequest {
	return &AlibabaLegalSuitJudgementGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitJudgementGetAPIRequest) Reset() {
	r._suitId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件id
func (r *AlibabaLegalSuitJudgementGetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabaLegalSuitJudgementGetAPIRequest) GetSuitId() int64 {
	return r._suitId
}

var poolAlibabaLegalSuitJudgementGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitJudgementGetRequest()
	},
}

// GetAlibabaLegalSuitJudgementGetRequest 从 sync.Pool 获取 AlibabaLegalSuitJudgementGetAPIRequest
func GetAlibabaLegalSuitJudgementGetAPIRequest() *AlibabaLegalSuitJudgementGetAPIRequest {
	return poolAlibabaLegalSuitJudgementGetAPIRequest.Get().(*AlibabaLegalSuitJudgementGetAPIRequest)
}

// ReleaseAlibabaLegalSuitJudgementGetAPIRequest 将 AlibabaLegalSuitJudgementGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitJudgementGetAPIRequest(v *AlibabaLegalSuitJudgementGetAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitJudgementGetAPIRequest.Put(v)
}
