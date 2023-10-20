package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Taobaoitemcatpropsmodificationget 查询商品类目属性变更
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
func Taobaoitemcatpropsmodificationget(clt *core.SDKClient, req *category.TaobaoitemcatpropsmodificationgetAPIRequest, session string) (*category.TaobaoitemcatpropsmodificationgetAPIResponse, error) {
	var resp category.TaobaoitemcatpropsmodificationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
