package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberidentityrescindfinish 取消确认
// alibaba.member.identity.rescindfinish
//
// 取消确认
func Alibabamemberidentityrescindfinish(clt *core.SDKClient, req *alimember.AlibabamemberidentityrescindfinishAPIRequest, session string) (*alimember.AlibabamemberidentityrescindfinishAPIResponse, error) {
	var resp alimember.AlibabamemberidentityrescindfinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
