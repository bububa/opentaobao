package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjoccalldispatcher 呼叫运力
// alibaba.mj.oc.calldispatcher
//
// 定时达呼叫运力接口
func Alibabamjoccalldispatcher(clt *core.SDKClient, req *mos.AlibabamjoccalldispatcherAPIRequest, session string) (*mos.AlibabamjoccalldispatcherAPIResponse, error) {
	var resp mos.AlibabamjoccalldispatcherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
