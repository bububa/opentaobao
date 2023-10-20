package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsQscoreGet 取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分
// taobao.simba.keywords.qscore.get
//
// 取得一个推广组的所有关键词的质量得分列表
func TaobaoSimbaKeywordsQscoreGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsQscoreGetAPIRequest, session string) (*simba.TaobaoSimbaKeywordsQscoreGetAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordsQscoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
