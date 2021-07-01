package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

/* TaobaoItempropsGet
获取标准商品类目属性
taobao.itemprops.get

通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。 */
func TaobaoItempropsGet(clt *core.SDKClient, req *category.TaobaoItempropsGetAPIRequest, session string) (*category.TaobaoItempropsGetAPIResponse, error) {
	var resp category.TaobaoItempropsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
