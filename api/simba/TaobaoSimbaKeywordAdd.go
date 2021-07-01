package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaKeywordAdd
（新）关键词新增接口
taobao.simba.keyword.add

（新）关键词更新相关接口 */
func TaobaoSimbaKeywordAdd(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordAddAPIRequest, session string) (*simba.TaobaoSimbaKeywordAddAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
