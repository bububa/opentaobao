package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlespuregistermodify 服务商spu挂载接口
// alibaba.idle.spu.register.modify
//
// 闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
func Alibabaidlespuregistermodify(clt *core.SDKClient, req *idle.AlibabaidlespuregistermodifyAPIRequest, session string) (*idle.AlibabaidlespuregistermodifyAPIResponse, error) {
	var resp idle.AlibabaidlespuregistermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
