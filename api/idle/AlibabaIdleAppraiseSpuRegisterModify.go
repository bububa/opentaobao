package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleappraisespuregistermodify 验货宝服务商spu挂载
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
func Alibabaidleappraisespuregistermodify(clt *core.SDKClient, req *idle.AlibabaidleappraisespuregistermodifyAPIRequest, session string) (*idle.AlibabaidleappraisespuregistermodifyAPIResponse, error) {
	var resp idle.AlibabaidleappraisespuregistermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
