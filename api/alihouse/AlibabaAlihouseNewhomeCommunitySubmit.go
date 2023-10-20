package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomecommunitysubmit 提交小区信息
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
func Alibabaalihousenewhomecommunitysubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomecommunitysubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomecommunitysubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomecommunitysubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
