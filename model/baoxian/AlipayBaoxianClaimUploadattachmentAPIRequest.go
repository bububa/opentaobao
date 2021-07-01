package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlipayBaoxianClaimUploadattachmentAPIRequest
资料上传接口 API请求
alipay.baoxian.claim.uploadattachment

给合作伙伴上传申请理赔材料 */
type AlipayBaoxianClaimUploadattachmentAPIRequest struct {
	model.Params
	// 外部业务号，唯一
	_outBizNo string
	// 业务来源
	_bizSource string
	// 标准产品ID
	_spNo string
	// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
	_attachmentKey string
	// 文件字节数组
	_attachmentByte *model.File
	// 是否base格式的字节数组
	_base64Bytes bool
	// 保单外部业务单号
	_policyBizNo string
	// 上传者用户标识
	_uploadUser string
}

// New
