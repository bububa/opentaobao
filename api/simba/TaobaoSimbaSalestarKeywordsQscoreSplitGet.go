package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarkeywordsqscoresplitget (新)销量明星质量分相关接口
// taobao.simba.salestar.keywords.qscore.split.get
//
// 获取关键词新的质量分
func Taobaosimbasalestarkeywordsqscoresplitget(clt *core.SDKClient, req *simba.TaobaosimbasalestarkeywordsqscoresplitgetAPIRequest, session string) (*simba.TaobaosimbasalestarkeywordsqscoresplitgetAPIResponse, error) {
	var resp simba.TaobaosimbasalestarkeywordsqscoresplitgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
