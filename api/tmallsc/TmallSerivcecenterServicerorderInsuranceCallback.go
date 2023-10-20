package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallSerivcecenterServicerorderInsuranceCallback 服务商回传保单信息
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
func TmallSerivcecenterServicerorderInsuranceCallback(clt *core.SDKClient, req *tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest, resp *tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
