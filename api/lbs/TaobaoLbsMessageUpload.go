package lbs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lbs"
)

// TaobaoLbsMessageUpload lbs数据采集
// taobao.lbs.message.upload
//
// lbs数据采集
func TaobaoLbsMessageUpload(clt *core.SDKClient, req *lbs.TaobaoLbsMessageUploadAPIRequest, resp *lbs.TaobaoLbsMessageUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
