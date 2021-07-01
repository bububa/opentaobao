package nlp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNlpSimilarityAPIRequest
文本语言相似度 API请求
taobao.nlp.similarity

文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。 */
type TaobaoNlpSimilarityAPIRequest struct {
	model.Params
	// 多文本内容
	_texts *Texts
}

// NewTaobaoNlpSimilarityRequest 初始化TaobaoNlpSimilarityAPIRequest对象
func NewTaobaoNlpSimilarityRequest() *TaobaoNlpSimilarityAPIRequest {
	return &TaobaoNlpSimilarityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNlpSimilarityAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.similarity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNlpSimilarityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Texts Setter
// 多文本内容
func (r *TaobaoNlpSimilarityAPIRequest) SetTexts(_texts *Texts) error {
	r._texts = _texts
	r.Set("texts", _texts)
	return nil
}

// Get Texts Getter
func (r TaobaoNlpSimilarityAPIRequest) GetTexts() *Texts {
	return r._texts
}
