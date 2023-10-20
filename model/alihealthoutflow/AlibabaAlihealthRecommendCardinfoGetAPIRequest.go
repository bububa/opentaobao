package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRecommendCardinfoGetAPIRequest 快应用卡片信息 API请求
// alibaba.alihealth.recommend.cardinfo.get
//
// 快应用卡片信息
type AlibabaAlihealthRecommendCardinfoGetAPIRequest struct {
	model.Params
	// 请求入参
	_cardRequest *QuickAppRequest
}

// NewAlibabaAlihealthRecommendCardinfoGetRequest 初始化AlibabaAlihealthRecommendCardinfoGetAPIRequest对象
func NewAlibabaAlihealthRecommendCardinfoGetRequest() *AlibabaAlihealthRecommendCardinfoGetAPIRequest {
	return &AlibabaAlihealthRecommendCardinfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthRecommendCardinfoGetAPIRequest) Reset() {
	r._cardRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRecommendCardinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.recommend.cardinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRecommendCardinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthRecommendCardinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardRequest is CardRequest Setter
// 请求入参
func (r *AlibabaAlihealthRecommendCardinfoGetAPIRequest) SetCardRequest(_cardRequest *QuickAppRequest) error {
	r._cardRequest = _cardRequest
	r.Set("card_request", _cardRequest)
	return nil
}

// GetCardRequest CardRequest Getter
func (r AlibabaAlihealthRecommendCardinfoGetAPIRequest) GetCardRequest() *QuickAppRequest {
	return r._cardRequest
}

var poolAlibabaAlihealthRecommendCardinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthRecommendCardinfoGetRequest()
	},
}

// GetAlibabaAlihealthRecommendCardinfoGetRequest 从 sync.Pool 获取 AlibabaAlihealthRecommendCardinfoGetAPIRequest
func GetAlibabaAlihealthRecommendCardinfoGetAPIRequest() *AlibabaAlihealthRecommendCardinfoGetAPIRequest {
	return poolAlibabaAlihealthRecommendCardinfoGetAPIRequest.Get().(*AlibabaAlihealthRecommendCardinfoGetAPIRequest)
}

// ReleaseAlibabaAlihealthRecommendCardinfoGetAPIRequest 将 AlibabaAlihealthRecommendCardinfoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthRecommendCardinfoGetAPIRequest(v *AlibabaAlihealthRecommendCardinfoGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthRecommendCardinfoGetAPIRequest.Put(v)
}
