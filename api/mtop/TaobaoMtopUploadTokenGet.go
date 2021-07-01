package mtop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtop"
)

/* TaobaoMtopUploadTokenGet
获取文件上传授权
taobao.mtop.upload.token.get

获取mtop文件上传授权 */
func TaobaoMtopUploadTokenGet(clt *core.SDKClient, req *mtop.TaobaoMtopUploadTokenGetAPIRequest, session string) (*mtop.TaobaoMtopUploadTokenGetAPIResponse, error) {
	var resp mtop.TaobaoMtopUploadTokenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
