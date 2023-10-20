package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemQuery 查询后端商品
// taobao.scitem.query
//
// 查询后端商品
func TaobaoScitemQuery(clt *core.SDKClient, req *fenxiao.TaobaoScitemQueryAPIRequest, resp *fenxiao.TaobaoScitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
