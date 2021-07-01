package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

/* TaobaoLarkPosBasedataGetworkstation
根据影城id工作站和macId获取工作站
taobao.lark.pos.basedata.getworkstation

获取单独工作站 */
func TaobaoLarkPosBasedataGetworkstation(clt *core.SDKClient, req *larkiot.TaobaoLarkPosBasedataGetworkstationAPIRequest, session string) (*larkiot.TaobaoLarkPosBasedataGetworkstationAPIResponse, error) {
	var resp larkiot.TaobaoLarkPosBasedataGetworkstationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
