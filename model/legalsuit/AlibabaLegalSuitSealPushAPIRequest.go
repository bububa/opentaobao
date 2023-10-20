package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitSealPushAPIRequest 法宝推送用印 API请求
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
type AlibabaLegalSuitSealPushAPIRequest struct {
	model.Params
	// 最外层list
	_sealTaskModel *SealTaskModel
}

// NewAlibabaLegalSuitSealPushRequest 初始化AlibabaLegalSuitSealPushAPIRequest对象
func NewAlibabaLegalSuitSealPushRequest() *AlibabaLegalSuitSealPushAPIRequest {
	return &AlibabaLegalSuitSealPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitSealPushAPIRequest) Reset() {
	r._sealTaskModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitSealPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.seal.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitSealPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitSealPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSealTaskModel is SealTaskModel Setter
// 最外层list
func (r *AlibabaLegalSuitSealPushAPIRequest) SetSealTaskModel(_sealTaskModel *SealTaskModel) error {
	r._sealTaskModel = _sealTaskModel
	r.Set("seal_task_model", _sealTaskModel)
	return nil
}

// GetSealTaskModel SealTaskModel Getter
func (r AlibabaLegalSuitSealPushAPIRequest) GetSealTaskModel() *SealTaskModel {
	return r._sealTaskModel
}

var poolAlibabaLegalSuitSealPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitSealPushRequest()
	},
}

// GetAlibabaLegalSuitSealPushRequest 从 sync.Pool 获取 AlibabaLegalSuitSealPushAPIRequest
func GetAlibabaLegalSuitSealPushAPIRequest() *AlibabaLegalSuitSealPushAPIRequest {
	return poolAlibabaLegalSuitSealPushAPIRequest.Get().(*AlibabaLegalSuitSealPushAPIRequest)
}

// ReleaseAlibabaLegalSuitSealPushAPIRequest 将 AlibabaLegalSuitSealPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitSealPushAPIRequest(v *AlibabaLegalSuitSealPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitSealPushAPIRequest.Put(v)
}
