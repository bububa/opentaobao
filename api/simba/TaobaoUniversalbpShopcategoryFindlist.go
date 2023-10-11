package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpShopcategoryFindlist 人群相关类目查询
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
func TaobaoUniversalbpShopcategoryFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpShopcategoryFindlistAPIRequest, session string) (*simba.TaobaoUniversalbpShopcategoryFindlistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpShopcategoryFindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
