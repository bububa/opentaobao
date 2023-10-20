package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpShopcategoryFindlist 人群相关类目查询
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
func TaobaoUniversalbpShopcategoryFindlist(clt *core.SDKClient, req *simba.TaobaoUniversalbpShopcategoryFindlistAPIRequest, resp *simba.TaobaoUniversalbpShopcategoryFindlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
