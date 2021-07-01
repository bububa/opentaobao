package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

/* TaobaoItemcatsGet
获取后台供卖家发布商品的标准商品类目
taobao.itemcats.get

获取后台供卖家发布商品的标准商品类目。 */
func TaobaoItemcatsGet(clt *core.SDKClient, req *category.TaobaoItemcatsGetAPIRequest, session string) (*category.TaobaoItemcatsGetAPIResponse, error) {
	var resp category.TaobaoItemcatsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
