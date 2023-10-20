package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderbtobitemupload 暗拍发布/编辑B2B商品
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
func Alibabaidletenderbtobitemupload(clt *core.SDKClient, req *idle.AlibabaidletenderbtobitemuploadAPIRequest, session string) (*idle.AlibabaidletenderbtobitemuploadAPIResponse, error) {
	var resp idle.AlibabaidletenderbtobitemuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
