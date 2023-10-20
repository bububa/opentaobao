package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomebusinesssync 商圈数据同步
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
func Alibabaalihousenewhomebusinesssync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomebusinesssyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomebusinesssyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomebusinesssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
