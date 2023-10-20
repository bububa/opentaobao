package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomestoreextendssync 门店扩展信息变更
// alibaba.alihouse.existinghome.store.extends.sync
//
// 门店扩展信息变更
func Alibabaalihouseexistinghomestoreextendssync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomestoreextendssyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomestoreextendssyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomestoreextendssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
