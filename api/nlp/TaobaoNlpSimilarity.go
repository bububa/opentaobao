package nlp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlp"
)

// TaobaoNlpSimilarity 文本语言相似度
// taobao.nlp.similarity
//
// 文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
func TaobaoNlpSimilarity(ctx context.Context, clt *core.SDKClient, req *nlp.TaobaoNlpSimilarityAPIRequest, resp *nlp.TaobaoNlpSimilarityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
