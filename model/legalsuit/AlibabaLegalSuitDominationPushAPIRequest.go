package legalsuit

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitDominationPushAPIRequest) Reset() {
	r._dominationDissentModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitDominationPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.domination.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitDominationPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitDominationPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDominationDissentModel is DominationDissentModel Setter
// 管辖信息
func (r *AlibabaLegalSuitDominationPushAPIRequest) SetDominationDissentModel(_dominationDissentModel *DominationDissentModel) error {
	r._dominationDissentModel = _dominationDissentModel
	r.Set("domination_dissent_model", _dominationDissentModel)
	return nil
}

// GetDominationDissentModel DominationDissentModel Getter
func (r AlibabaLegalSuitDominationPushAPIRequest) GetDominationDissentModel() *DominationDissentModel {
	return r._dominationDissentModel
}

var poolAlibabaLegalSuitDominationPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitDominationPushRequest()
	},
}

// GetAlibabaLegalSuitDominationPushRequest 从 sync.Pool 获取 AlibabaLegalSuitDominationPushAPIRequest
func GetAlibabaLegalSuitDominationPushAPIRequest() *AlibabaLegalSuitDominationPushAPIRequest {
	return poolAlibabaLegalSuitDominationPushAPIRequest.Get().(*AlibabaLegalSuitDominationPushAPIRequest)
}

// ReleaseAlibabaLegalSuitDominationPushAPIRequest 将 AlibabaLegalSuitDominationPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitDominationPushAPIRequest(v *AlibabaLegalSuitDominationPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitDominationPushAPIRequest.Put(v)
}
