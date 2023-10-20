package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomecommunitybrokersubmit 提交小区专家
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
func Alibabaalihouseexistinghomecommunitybrokersubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomecommunitybrokersubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomecommunitybrokersubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomecommunitybrokersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
