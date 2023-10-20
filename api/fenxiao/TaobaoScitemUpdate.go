package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemUpdate 根据商品ID或商家编码修改后端商品
// taobao.scitem.update
//
// 根据商品ID或商家编码修改后端商品
func TaobaoScitemUpdate(clt *core.SDKClient, req *fenxiao.TaobaoScitemUpdateAPIRequest, resp *fenxiao.TaobaoScitemUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
