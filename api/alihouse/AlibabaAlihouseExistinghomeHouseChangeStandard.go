package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousechangestandard 委托房源变更标准房源
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
func Alibabaalihouseexistinghomehousechangestandard(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousechangestandardAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousechangestandardAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousechangestandardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
