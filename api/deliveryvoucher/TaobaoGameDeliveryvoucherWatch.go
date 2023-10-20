package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// Taobaogamedeliveryvoucherwatch 监控预约数据
// taobao.game.deliveryvoucher.watch
//
// 监控预约数据
func Taobaogamedeliveryvoucherwatch(clt *core.SDKClient, req *deliveryvoucher.TaobaogamedeliveryvoucherwatchAPIRequest, session string) (*deliveryvoucher.TaobaogamedeliveryvoucherwatchAPIResponse, error) {
	var resp deliveryvoucher.TaobaogamedeliveryvoucherwatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
