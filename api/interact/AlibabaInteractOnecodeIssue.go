package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractonecodeissue onecode代码通用鉴权
// alibaba.interact.onecode.issue
//
// 手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
func Alibabainteractonecodeissue(clt *core.SDKClient, req *interact.AlibabainteractonecodeissueAPIRequest, session string) (*interact.AlibabainteractonecodeissueAPIResponse, error) {
	var resp interact.AlibabainteractonecodeissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
