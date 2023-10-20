package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalSuitFileUpload 诉讼文件上传接口
// alibaba.legal.suit.file.upload
//
// 上传文件接口
func AlibabaLegalSuitFileUpload(clt *core.SDKClient, req *legalcase.AlibabaLegalSuitFileUploadAPIRequest, resp *legalcase.AlibabaLegalSuitFileUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
