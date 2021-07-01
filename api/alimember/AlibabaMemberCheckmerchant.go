package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

/* AlibabaMemberCheckmerchant
校验商家身份
alibaba.member.checkmerchant

校验商家身份 */
func AlibabaMemberCheckmerchant(clt *core.SDKClient, req *alimember.AlibabaMemberCheckmerchantAPIRequest, session string) (*alimember.AlibabaMemberCheckmerchantAPIResponse, error) {
	var resp alimember.AlibabaMemberCheckmerchantAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
