package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthZyUploadrelationAPIRequest
中药片关联关系上传 API请求
alibaba.alihealth.zy.uploadrelation

中药片关联关系上传 */
type AlibabaAlihealthZyUploadrelationAPIRequest struct {
	model.Params
	// 关联关系文件信息
	_saveCodeRelation *SaveCodeRelationType
	// affirmFlag
	_affirmFlag string
	// fileContent
	_fileContent string
	// 加密之后的上传的关联关系文件内容
	_fileContentString string
	// 企业ID
	_refEntId string
	// 客户端类型
	_clientType string
}

// New
