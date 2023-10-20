package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjoconlineticketnoget 线上小票号获取
// alibaba.mj.oc.online.ticketno.get
//
// 线上小票号获取
func Alibabamjoconlineticketnoget(clt *core.SDKClient, req *mos.AlibabamjoconlineticketnogetAPIRequest, session string) (*mos.AlibabamjoconlineticketnogetAPIResponse, error) {
	var resp mos.AlibabamjoconlineticketnogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
