package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvmediaupload 闲鱼服务商-图片上传
// alibaba.idle.isv.media.upload
//
// 供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
func Alibabaidleisvmediaupload(clt *core.SDKClient, req *idleisv.AlibabaidleisvmediauploadAPIRequest, session string) (*idleisv.AlibabaidleisvmediauploadAPIResponse, error) {
	var resp idleisv.AlibabaidleisvmediauploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
