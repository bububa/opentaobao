package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopdepartsync 外部部门同步
// alitrip.btrip.corpop.depart.sync
//
// 同步外部平台部门信息至商旅内部
func Alitripbtripcorpopdepartsync(clt *core.SDKClient, req *btrip.AlitripbtripcorpopdepartsyncAPIRequest, session string) (*btrip.AlitripbtripcorpopdepartsyncAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopdepartsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
