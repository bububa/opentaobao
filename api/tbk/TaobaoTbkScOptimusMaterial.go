package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScOptimusMaterial 淘宝客-服务商-物料精选
// taobao.tbk.sc.optimus.material
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广者对应的推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
func TaobaoTbkScOptimusMaterial(clt *core.SDKClient, req *tbk.TaobaoTbkScOptimusMaterialAPIRequest, resp *tbk.TaobaoTbkScOptimusMaterialAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
