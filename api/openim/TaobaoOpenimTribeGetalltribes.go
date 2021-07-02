package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeGetalltribes 获取用户群列表
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
func TaobaoOpenimTribeGetalltribes(clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetalltribesAPIRequest, session string) (*openim.TaobaoOpenimTribeGetalltribesAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeGetalltribesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
