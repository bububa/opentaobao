package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsqscoresplitget 新质量分服务
// taobao.simba.keywords.qscore.split.get
//
// 获取关键词新的质量分
func Taobaosimbakeywordsqscoresplitget(clt *core.SDKClient, req *simba.TaobaosimbakeywordsqscoresplitgetAPIRequest, session string) (*simba.TaobaosimbakeywordsqscoresplitgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsqscoresplitgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
