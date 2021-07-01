package mtop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMtopUploadTokenGetAPIRequest
获取文件上传授权 API请求
taobao.mtop.upload.token.get

获取mtop文件上传授权 */
type TaobaoMtopUploadTokenGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramUploadTokenRequest *UploadTokenRequestV
}

// New
