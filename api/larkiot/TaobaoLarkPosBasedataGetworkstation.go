package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// TaobaoLarkPosBasedataGetworkstation 根据影城id工作站和macId获取工作站
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
func TaobaoLarkPosBasedataGetworkstation(clt *core.SDKClient, req *larkiot.TaobaoLarkPosBasedataGetworkstationAPIRequest, resp *larkiot.TaobaoLarkPosBasedataGetworkstationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
