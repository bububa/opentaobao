package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemUpload 暗拍发布/编辑B2B商品
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
func AlibabaIdleTenderBtobItemUpload(clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemUploadAPIRequest, resp *idle.AlibabaIdleTenderBtobItemUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
