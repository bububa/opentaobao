package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleadvmaterialupload 闲鱼用户增长素材中心素材上传接口
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
func Alibabaidleadvmaterialupload(clt *core.SDKClient, req *idle.AlibabaidleadvmaterialuploadAPIRequest, session string) (*idle.AlibabaidleadvmaterialuploadAPIResponse, error) {
	var resp idle.AlibabaidleadvmaterialuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
