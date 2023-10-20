package nlp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpSimilarityAPIRequest 文本语言相似度 API请求
// taobao.nlp.similarity
//
// 文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
type TaobaoNlpSimilarityAPIRequest struct {
	model.Params
	// 多文本内容
	_texts *Texts
}

// NewTaobaoNlpSimilarityRequest 初始化TaobaoNlpSimilarityAPIRequest对象
func NewTaobaoNlpSimilarityRequest() *TaobaoNlpSimilarityAPIRequest {
	return &TaobaoNlpSimilarityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoNlpSimilarityAPIRequest) Reset() {
	r._texts = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNlpSimilarityAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.similarity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNlpSimilarityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoNlpSimilarityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTexts is Texts Setter
// 多文本内容
func (r *TaobaoNlpSimilarityAPIRequest) SetTexts(_texts *Texts) error {
	r._texts = _texts
	r.Set("texts", _texts)
	return nil
}

// GetTexts Texts Getter
func (r TaobaoNlpSimilarityAPIRequest) GetTexts() *Texts {
	return r._texts
}

var poolTaobaoNlpSimilarityAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoNlpSimilarityRequest()
	},
}

// GetTaobaoNlpSimilarityRequest 从 sync.Pool 获取 TaobaoNlpSimilarityAPIRequest
func GetTaobaoNlpSimilarityAPIRequest() *TaobaoNlpSimilarityAPIRequest {
	return poolTaobaoNlpSimilarityAPIRequest.Get().(*TaobaoNlpSimilarityAPIRequest)
}

// ReleaseTaobaoNlpSimilarityAPIRequest 将 TaobaoNlpSimilarityAPIRequest 放入 sync.Pool
func ReleaseTaobaoNlpSimilarityAPIRequest(v *TaobaoNlpSimilarityAPIRequest) {
	v.Reset()
	poolTaobaoNlpSimilarityAPIRequest.Put(v)
}
