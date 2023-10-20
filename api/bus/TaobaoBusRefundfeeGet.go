package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusrefundfeeget 查询退票费用明细
// taobao.bus.refundfee.get
//
// 查询退票的费用信息
func Taobaobusrefundfeeget(clt *core.SDKClient, req *bus.TaobaobusrefundfeegetAPIRequest, session string) (*bus.TaobaobusrefundfeegetAPIResponse, error) {
	var resp bus.TaobaobusrefundfeegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
