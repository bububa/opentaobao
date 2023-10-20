package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomenewreviewsync 新测评基础信息接口
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
func Alibabaalihousenewhomenewreviewsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomenewreviewsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomenewreviewsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomenewreviewsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
