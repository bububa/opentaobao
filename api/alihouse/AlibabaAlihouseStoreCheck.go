package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousestorecheck 门店对账查询工具
// alibaba.alihouse.store.check
//
// 门店对账查询工具
func Alibabaalihousestorecheck(clt *core.SDKClient, req *alihouse.AlibabaalihousestorecheckAPIRequest, session string) (*alihouse.AlibabaalihousestorecheckAPIResponse, error) {
	var resp alihouse.AlibabaalihousestorecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
