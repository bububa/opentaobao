package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeposapplysubmit pos申请接口
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
func Alibabaalihouseexistinghomeposapplysubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeposapplysubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeposapplysubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeposapplysubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
