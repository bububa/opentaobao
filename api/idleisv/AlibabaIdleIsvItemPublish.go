package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvitempublish 服务商闲鱼商品发布
// alibaba.idle.isv.item.publish
//
// 服务商ISV闲鱼商品发布
func Alibabaidleisvitempublish(clt *core.SDKClient, req *idleisv.AlibabaidleisvitempublishAPIRequest, session string) (*idleisv.AlibabaidleisvitempublishAPIResponse, error) {
	var resp idleisv.AlibabaidleisvitempublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
