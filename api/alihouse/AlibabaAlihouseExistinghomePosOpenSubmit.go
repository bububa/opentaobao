package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeposopensubmit pos进件接口
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
func Alibabaalihouseexistinghomeposopensubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeposopensubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeposopensubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeposopensubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
