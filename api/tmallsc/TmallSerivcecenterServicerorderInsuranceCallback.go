package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallSerivcecenterServicerorderInsuranceCallback 服务商回传保单信息
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
func TmallSerivcecenterServicerorderInsuranceCallback(clt *core.SDKClient, req *tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest, session string) (*tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse, error) {
	var resp tmallsc.TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
