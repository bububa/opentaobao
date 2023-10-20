package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomesupportsync 周边配套数据同步
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
func Alibabaalihousenewhomesupportsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomesupportsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomesupportsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomesupportsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
