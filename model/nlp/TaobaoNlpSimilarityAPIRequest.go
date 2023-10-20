package nlp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaonlpsimilarityAPIRequest 文本语言相似度 API请求
// taobao.nlp.similarity
//
// 文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
type TaobaonlpsimilarityAPIRequest struct {
	model.Params
	// 多文本内容
	_texts *Texts
}

// NewTaobaonlpsimilarityRequest 初始化TaobaonlpsimilarityAPIRequest对象
func NewTaobaonlpsimilarityRequest() *TaobaonlpsimilarityAPIRequest {
	return &TaobaonlpsimilarityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaonlpsimilarityAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.similarity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaonlpsimilarityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaonlpsimilarityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTexts is Texts Setter
// 多文本内容
func (r *TaobaonlpsimilarityAPIRequest) SetTexts(_texts *Texts) error {
	r._texts = _texts
	r.Set("texts", _texts)
	return nil
}

// GetTexts Texts Getter
func (r TaobaonlpsimilarityAPIRequest) GetTexts() *Texts {
	return r._texts
}
