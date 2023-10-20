package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomebaselabelsubmit 提交标签库
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
func Alibabaalihousenewhomebaselabelsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomebaselabelsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomebaselabelsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomebaselabelsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
