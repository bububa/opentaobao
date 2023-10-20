package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationgetinfouuid 根据uuid用户抽奖次数限制
// taobao.degoperation.get.info.uuid
//
// 根据uuid用户抽奖次数限制
func Taobaodegoperationgetinfouuid(clt *core.SDKClient, req *degoperation.TaobaodegoperationgetinfouuidAPIRequest, session string) (*degoperation.TaobaodegoperationgetinfouuidAPIResponse, error) {
	var resp degoperation.TaobaodegoperationgetinfouuidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
