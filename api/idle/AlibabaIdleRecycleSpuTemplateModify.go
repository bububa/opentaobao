package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerecyclesputemplatemodify 闲鱼接收回收商spu模板挂载信息
// alibaba.idle.recycle.spu.template.modify
//
// 闲鱼接收回收商spu模板挂载信息
func Alibabaidlerecyclesputemplatemodify(clt *core.SDKClient, req *idle.AlibabaidlerecyclesputemplatemodifyAPIRequest, session string) (*idle.AlibabaidlerecyclesputemplatemodifyAPIResponse, error) {
	var resp idle.AlibabaidlerecyclesputemplatemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
