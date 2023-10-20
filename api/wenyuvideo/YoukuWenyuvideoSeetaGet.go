package wenyuvideo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// YoukuWenyuvideoSeetaGet 只看TA
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
func YoukuWenyuvideoSeetaGet(clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoSeetaGetAPIRequest, resp *wenyuvideo.YoukuWenyuvideoSeetaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
