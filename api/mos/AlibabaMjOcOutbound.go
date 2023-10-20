package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocoutbound 零售商品发货
// alibaba.mj.oc.outbound
//
// 用于接收发货的数据
func Alibabamjocoutbound(clt *core.SDKClient, req *mos.AlibabamjocoutboundAPIRequest, session string) (*mos.AlibabamjocoutboundAPIResponse, error) {
	var resp mos.AlibabamjocoutboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
