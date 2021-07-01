package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

/* AlibabaBenefitSend
发奖接口
alibaba.benefit.send

发奖接口 */
func AlibabaBenefitSend(clt *core.SDKClient, req *user.AlibabaBenefitSendAPIRequest, session string) (*user.AlibabaBenefitSendAPIResponse, error) {
	var resp user.AlibabaBenefitSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
