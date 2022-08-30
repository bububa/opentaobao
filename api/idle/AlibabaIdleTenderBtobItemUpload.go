package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemUpload 暗拍发布/编辑B2B商品
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
func AlibabaIdleTenderBtobItemUpload(clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemUploadAPIRequest, session string) (*idle.AlibabaIdleTenderBtobItemUploadAPIResponse, error) {
	var resp idle.AlibabaIdleTenderBtobItemUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
