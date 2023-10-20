package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkActivityInfoGet 淘宝客-推广者-官方活动转链
// taobao.tbk.activity.info.get
//
// 支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
func TaobaoTbkActivityInfoGet(clt *core.SDKClient, req *tbk.TaobaoTbkActivityInfoGetAPIRequest, session string) (*tbk.TaobaoTbkActivityInfoGetAPIResponse, error) {
	var resp tbk.TaobaoTbkActivityInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
