package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribegetalltribes 获取用户群列表
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
func Taobaoopenimtribegetalltribes(clt *core.SDKClient, req *openim.TaobaoopenimtribegetalltribesAPIRequest, session string) (*openim.TaobaoopenimtribegetalltribesAPIResponse, error) {
	var resp openim.TaobaoopenimtribegetalltribesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
