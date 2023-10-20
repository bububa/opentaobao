package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusnumbersupdate 汽车票车次更新服务
// taobao.bus.numbers.update
//
// 用于汽车票车次信息的新增、更新和逻辑删除
func Taobaobusnumbersupdate(clt *core.SDKClient, req *bus.TaobaobusnumbersupdateAPIRequest, session string) (*bus.TaobaobusnumbersupdateAPIResponse, error) {
	var resp bus.TaobaobusnumbersupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
