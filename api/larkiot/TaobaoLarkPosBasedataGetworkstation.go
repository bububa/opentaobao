package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// Taobaolarkposbasedatagetworkstation 根据影城id工作站和macId获取工作站
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
func Taobaolarkposbasedatagetworkstation(clt *core.SDKClient, req *larkiot.TaobaolarkposbasedatagetworkstationAPIRequest, session string) (*larkiot.TaobaolarkposbasedatagetworkstationAPIResponse, error) {
	var resp larkiot.TaobaolarkposbasedatagetworkstationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
