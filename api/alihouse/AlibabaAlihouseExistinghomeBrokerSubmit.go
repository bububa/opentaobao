package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrokersubmit 提交经纪人信息
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
func Alibabaalihouseexistinghomebrokersubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrokersubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrokersubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrokersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
