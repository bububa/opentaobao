package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyorderdesensitizephoneget 获取订单隐私号
// alibaba.aelophy.order.desensitizephone.get
//
// 获取订单隐私号
func Alibabaaelophyorderdesensitizephoneget(clt *core.SDKClient, req *wdk.AlibabaaelophyorderdesensitizephonegetAPIRequest, session string) (*wdk.AlibabaaelophyorderdesensitizephonegetAPIResponse, error) {
	var resp wdk.AlibabaaelophyorderdesensitizephonegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
