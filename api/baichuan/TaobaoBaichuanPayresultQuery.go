package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanPayresultQuery 百川支付完成回调
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
func TaobaoBaichuanPayresultQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanPayresultQueryAPIRequest, resp *baichuan.TaobaoBaichuanPayresultQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
