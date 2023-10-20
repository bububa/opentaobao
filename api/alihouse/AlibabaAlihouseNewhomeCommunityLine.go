package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomecommunityline 小区上下架
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
func Alibabaalihousenewhomecommunityline(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomecommunitylineAPIRequest, session string) (*alihouse.AlibabaalihousenewhomecommunitylineAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomecommunitylineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
