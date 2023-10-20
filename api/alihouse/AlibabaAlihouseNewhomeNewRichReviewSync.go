package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomenewrichreviewsync 新评测改造接口同步
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
func Alibabaalihousenewhomenewrichreviewsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomenewrichreviewsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomenewrichreviewsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomenewrichreviewsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
