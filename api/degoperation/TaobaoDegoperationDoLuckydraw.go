package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationDoLuckydraw 激励抽奖
// taobao.degoperation.do.luckydraw
//
// 激励平台抽奖接口。用户可以通过接口完成抽奖功能
func TaobaoDegoperationDoLuckydraw(clt *core.SDKClient, req *degoperation.TaobaoDegoperationDoLuckydrawAPIRequest, resp *degoperation.TaobaoDegoperationDoLuckydrawAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
