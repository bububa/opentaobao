package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsvonAdd 创建一批关键词
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
func TaobaoSimbaKeywordsvonAdd(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsvonAddAPIRequest, resp *simba.TaobaoSimbaKeywordsvonAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
