package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemAdd 发布后端商品
// taobao.scitem.add
//
// 发布后端商品
func TaobaoScitemAdd(clt *core.SDKClient, req *fenxiao.TaobaoScitemAddAPIRequest, resp *fenxiao.TaobaoScitemAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
