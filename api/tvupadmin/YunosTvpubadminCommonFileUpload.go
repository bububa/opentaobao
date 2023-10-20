package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminCommonFileUpload 文件上传API
// yunos.tvpubadmin.common.file.upload
//
// 文件上传服务
func YunosTvpubadminCommonFileUpload(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminCommonFileUploadAPIRequest, resp *tvupadmin.YunosTvpubadminCommonFileUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
