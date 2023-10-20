package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsbyadgroupidGet 取得一个推广组的所有关键词
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
func TaobaoSimbaKeywordsbyadgroupidGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsbyadgroupidGetAPIRequest, resp *simba.TaobaoSimbaKeywordsbyadgroupidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
