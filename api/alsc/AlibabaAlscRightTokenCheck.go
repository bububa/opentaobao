package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscRightTokenCheck 实物奖品凭证校验
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
func AlibabaAlscRightTokenCheck(clt *core.SDKClient, req *alsc.AlibabaAlscRightTokenCheckAPIRequest, session string) (*alsc.AlibabaAlscRightTokenCheckAPIResponse, error) {
	var resp alsc.AlibabaAlscRightTokenCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
