package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

/* AlibabaInteractMediaArtwork
原图相关鉴权接口
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口 */
func AlibabaInteractMediaArtwork(clt *core.SDKClient, req *mtopopen.AlibabaInteractMediaArtworkAPIRequest, session string) (*mtopopen.AlibabaInteractMediaArtworkAPIResponse, error) {
	var resp mtopopen.AlibabaInteractMediaArtworkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
