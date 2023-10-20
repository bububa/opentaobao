package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScActivityInfoGet 淘宝客-服务商-官方活动转链
// taobao.tbk.sc.activity.info.get
//
// 服务商使用。支持入参推广者对应的推广位和官方活动会场ID，获取活动信息和推广者的推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
func TaobaoTbkScActivityInfoGet(clt *core.SDKClient, req *tbk.TaobaoTbkScActivityInfoGetAPIRequest, resp *tbk.TaobaoTbkScActivityInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
