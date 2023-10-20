package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelUpdate 酒店更新接口（ID不存在自动新增）
// taobao.xhotel.update
//
// 酒店更新接口
func TaobaoXhotelUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
