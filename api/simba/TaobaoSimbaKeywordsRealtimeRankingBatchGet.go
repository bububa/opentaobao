package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsrealtimerankingbatchget 获取关键词的新版实时排名
// taobao.simba.keywords.realtime.ranking.batch.get
//
// 根据关键词ID获取关键词的新版实时排名
func Taobaosimbakeywordsrealtimerankingbatchget(clt *core.SDKClient, req *simba.TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest, session string) (*simba.TaobaosimbakeywordsrealtimerankingbatchgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsrealtimerankingbatchgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
