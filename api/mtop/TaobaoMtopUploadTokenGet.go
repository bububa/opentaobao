package mtop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtop"
)

// TaobaoMtopUploadTokenGet 获取文件上传授权
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
func TaobaoMtopUploadTokenGet(clt *core.SDKClient, req *mtop.TaobaoMtopUploadTokenGetAPIRequest, resp *mtop.TaobaoMtopUploadTokenGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
