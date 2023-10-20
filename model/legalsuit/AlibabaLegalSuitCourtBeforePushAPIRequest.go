package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtBeforePushAPIRequest 更新或保存庭前信息 API请求
// alibaba.legal.suit.court.before.push
//
// 更新或者保存庭前信息
type AlibabaLegalSuitCourtBeforePushAPIRequest struct {
	model.Params
	// 庭前信息
	_beforeCourtModel *BeforeCourtModel
}

// NewAlibabaLegalSuitCourtBeforePushRequest 初始化AlibabaLegalSuitCourtBeforePushAPIRequest对象
func NewAlibabaLegalSuitCourtBeforePushRequest() *AlibabaLegalSuitCourtBeforePushAPIRequest {
	return &AlibabaLegalSuitCourtBeforePushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitCourtBeforePushAPIRequest) Reset() {
	r._beforeCourtModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.before.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBeforeCourtModel is BeforeCourtModel Setter
// 庭前信息
func (r *AlibabaLegalSuitCourtBeforePushAPIRequest) SetBeforeCourtModel(_beforeCourtModel *BeforeCourtModel) error {
	r._beforeCourtModel = _beforeCourtModel
	r.Set("before_court_model", _beforeCourtModel)
	return nil
}

// GetBeforeCourtModel BeforeCourtModel Getter
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetBeforeCourtModel() *BeforeCourtModel {
	return r._beforeCourtModel
}

var poolAlibabaLegalSuitCourtBeforePushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitCourtBeforePushRequest()
	},
}

// GetAlibabaLegalSuitCourtBeforePushRequest 从 sync.Pool 获取 AlibabaLegalSuitCourtBeforePushAPIRequest
func GetAlibabaLegalSuitCourtBeforePushAPIRequest() *AlibabaLegalSuitCourtBeforePushAPIRequest {
	return poolAlibabaLegalSuitCourtBeforePushAPIRequest.Get().(*AlibabaLegalSuitCourtBeforePushAPIRequest)
}

// ReleaseAlibabaLegalSuitCourtBeforePushAPIRequest 将 AlibabaLegalSuitCourtBeforePushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitCourtBeforePushAPIRequest(v *AlibabaLegalSuitCourtBeforePushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitCourtBeforePushAPIRequest.Put(v)
}
