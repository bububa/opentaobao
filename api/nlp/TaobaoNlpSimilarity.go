package nlp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlp"
)

// Taobaonlpsimilarity 文本语言相似度
// taobao.nlp.similarity
//
// 文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
func Taobaonlpsimilarity(clt *core.SDKClient, req *nlp.TaobaonlpsimilarityAPIRequest, session string) (*nlp.TaobaonlpsimilarityAPIResponse, error) {
	var resp nlp.TaobaonlpsimilarityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
