package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaKeywordsbykeywordidsGet
根据一个关键词Id列表取得一组关键词
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词 */
func TaobaoSimbaKeywordsbykeywordidsGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsbykeywordidsGetAPIRequest, session string) (*simba.TaobaoSimbaKeywordsbykeywordidsGetAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordsbykeywordidsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
