package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScAdzoneCreate 淘宝客-服务商-创建推广者位
// taobao.tbk.sc.adzone.create
//
// 提供淘宝客创建广告位
func TaobaoTbkScAdzoneCreate(clt *core.SDKClient, req *tbk.TaobaoTbkScAdzoneCreateAPIRequest, resp *tbk.TaobaoTbkScAdzoneCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
