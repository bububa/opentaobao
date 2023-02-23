package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItempropsGet 获取标准商品类目属性
// taobao.itemprops.get
//
// 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
func TaobaoItempropsGet(clt *core.SDKClient, req *tbitem.TaobaoItempropsGetAPIRequest, session string) (*tbitem.TaobaoItempropsGetAPIResponse, error) {
	var resp tbitem.TaobaoItempropsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
