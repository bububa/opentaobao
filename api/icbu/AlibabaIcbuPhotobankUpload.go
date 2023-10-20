package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankUpload 图片银行图片上传开放接口
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
func AlibabaIcbuPhotobankUpload(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankUploadAPIRequest, resp *icbu.AlibabaIcbuPhotobankUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
