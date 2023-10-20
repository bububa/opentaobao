package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomemetrosync 地铁数据同步
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
func Alibabaalihousenewhomemetrosync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomemetrosyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomemetrosyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomemetrosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
