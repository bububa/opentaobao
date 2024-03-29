package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemMapDelete 失效指定用户的商品与后端商品的映射关系
// taobao.scitem.map.delete
//
// 根据后端商品Id，失效指定用户的商品与后端商品的映射关系
func TaobaoScitemMapDelete(clt *core.SDKClient, req *fenxiao.TaobaoScitemMapDeleteAPIRequest, resp *fenxiao.TaobaoScitemMapDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
