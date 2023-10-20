package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslIteminfoPut 电子价签显示用商品信息写入
// taobao.uscesl.iteminfo.put
//
// 用于电子价签上显示的商品信息的写入，包含价格及促销信息
func TaobaoUsceslIteminfoPut(clt *core.SDKClient, req *uscesl.TaobaoUsceslIteminfoPutAPIRequest, resp *uscesl.TaobaoUsceslIteminfoPutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
