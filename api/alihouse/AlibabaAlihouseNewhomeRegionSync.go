package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeregionsync 城区数据同步
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
func Alibabaalihousenewhomeregionsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeregionsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeregionsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeregionsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
