package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiordercomplain 用户投诉
// alibaba.happytrip.taxi.order.complain
//
// 一个订单只能投诉一次，不可重复投诉
//
// 投诉选项
// 0	其他原因；
// 1	司机原因导致行程被取消；
// 2	服务态度恶劣；
// 3	未坐车产生费用；
// 4	额外收取不合理费用；
// 5	路不熟多产生费用；
// 6	提前计费；
// 7	未及时结束计费；
// 8	司机绕路；
// 9	司机迟到；
// 10	司机爽约或拒载；
// 11	骚扰乘客；
// 12	危险驾驶；
// 13	不是订单显示车辆或司机；
func Alibabahappytriptaxiordercomplain(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiordercomplainAPIRequest, session string) (*happytrip.AlibabahappytriptaxiordercomplainAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiordercomplainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
