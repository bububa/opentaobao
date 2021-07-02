package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// TaobaoItemCatpropsModificationGet 查询商品类目属性变更
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
func TaobaoItemCatpropsModificationGet(clt *core.SDKClient, req *category.TaobaoItemCatpropsModificationGetAPIRequest, session string) (*category.TaobaoItemCatpropsModificationGetAPIResponse, error) {
	var resp category.TaobaoItemCatpropsModificationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
