package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdCreate 淘宝客-服务商-创建人群标签
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
func TaobaoTbkScUcrowdCreate(clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdCreateAPIRequest, resp *tbk.TaobaoTbkScUcrowdCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
