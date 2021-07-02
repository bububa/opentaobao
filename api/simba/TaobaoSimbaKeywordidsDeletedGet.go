package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordidsDeletedGet 获取删除的词ID
// taobao.simba.keywordids.deleted.get
//
// 获取删除的词ID
func TaobaoSimbaKeywordidsDeletedGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordidsDeletedGetAPIRequest, session string) (*simba.TaobaoSimbaKeywordidsDeletedGetAPIResponse, error) {
	var resp simba.TaobaoSimbaKeywordidsDeletedGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
