package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousechangestoretype 融合店迁移门店
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
func Alibabaalihousechangestoretype(clt *core.SDKClient, req *alihouse.AlibabaalihousechangestoretypeAPIRequest, session string) (*alihouse.AlibabaalihousechangestoretypeAPIResponse, error) {
	var resp alihouse.AlibabaalihousechangestoretypeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
