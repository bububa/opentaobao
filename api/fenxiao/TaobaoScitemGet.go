package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemGet 根据id查询商品
// taobao.scitem.get
//
// 根据id查询商品
func TaobaoScitemGet(clt *core.SDKClient, req *fenxiao.TaobaoScitemGetAPIRequest, resp *fenxiao.TaobaoScitemGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
