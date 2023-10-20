package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjpresalesettlementaddlist 预售结算数据回传
// alibaba.mj.presale.settlement.addlist
//
// 用于预售活动结算数据的回传。
func Alibabamjpresalesettlementaddlist(clt *core.SDKClient, req *mos.AlibabamjpresalesettlementaddlistAPIRequest, session string) (*mos.AlibabamjpresalesettlementaddlistAPIResponse, error) {
	var resp mos.AlibabamjpresalesettlementaddlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
