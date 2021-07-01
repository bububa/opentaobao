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

// New
