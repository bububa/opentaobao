package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractmediaartwork 原图相关鉴权接口
// alibaba.interact.media.artwork
//
// 拍摄并上传原图相关鉴权接口
func Alibabainteractmediaartwork(clt *core.SDKClient, req *mtopopen.AlibabainteractmediaartworkAPIRequest, session string) (*mtopopen.AlibabainteractmediaartworkAPIResponse, error) {
	var resp mtopopen.AlibabainteractmediaartworkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
