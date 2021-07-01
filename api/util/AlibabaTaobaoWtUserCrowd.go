package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* AlibabaTaobaoWtUserCrowd
校验用户是否为新人
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型 */
func AlibabaTaobaoWtUserCrowd(clt *core.SDKClient, req *util.AlibabaTaobaoWtUserCrowdAPIRequest, session string) (*util.AlibabaTaobaoWtUserCrowdAPIResponse, error) {
	var resp util.AlibabaTaobaoWtUserCrowdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
