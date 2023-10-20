package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordscatqscoreget 取得一个推广组的所有关键词和类目出价的质量得分
// taobao.simba.keywordscat.qscore.get
//
// 取得一个推广组的所有关键词和类目出价的质量得分列表
func Taobaosimbakeywordscatqscoreget(clt *core.SDKClient, req *simba.TaobaosimbakeywordscatqscoregetAPIRequest, session string) (*simba.TaobaosimbakeywordscatqscoregetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordscatqscoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
