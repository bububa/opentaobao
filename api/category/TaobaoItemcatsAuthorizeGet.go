package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Taobaoitemcatsauthorizeget 查询商家被授权品牌列表和类目列表
// taobao.itemcats.authorize.get
//
// 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
func Taobaoitemcatsauthorizeget(clt *core.SDKClient, req *category.TaobaoitemcatsauthorizegetAPIRequest, session string) (*category.TaobaoitemcatsauthorizegetAPIResponse, error) {
	var resp category.TaobaoitemcatsauthorizegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
