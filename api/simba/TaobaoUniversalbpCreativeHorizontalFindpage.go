package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativeHorizontalFindpage 横向管理创意分页查询
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
func TaobaoUniversalbpCreativeHorizontalFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest, session string) (*simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
