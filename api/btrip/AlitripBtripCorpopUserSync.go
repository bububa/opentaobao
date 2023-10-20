package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopusersync 外部人员同步
// alitrip.btrip.corpop.user.sync
//
// 同步外部平台用户信息至商旅内部
func Alitripbtripcorpopusersync(clt *core.SDKClient, req *btrip.AlitripbtripcorpopusersyncAPIRequest, session string) (*btrip.AlitripbtripcorpopusersyncAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopusersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
