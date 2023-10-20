package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsvonadd 创建一批关键词
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
func Taobaosimbakeywordsvonadd(clt *core.SDKClient, req *simba.TaobaosimbakeywordsvonaddAPIRequest, session string) (*simba.TaobaosimbakeywordsvonaddAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsvonaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
