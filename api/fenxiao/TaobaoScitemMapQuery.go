package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemMapQuery 查找IC商品或分销商品与后端商品的关联信息
// taobao.scitem.map.query
//
// 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
func TaobaoScitemMapQuery(clt *core.SDKClient, req *fenxiao.TaobaoScitemMapQueryAPIRequest, resp *fenxiao.TaobaoScitemMapQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
