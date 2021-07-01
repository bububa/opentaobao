package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadrelationAPIRequest
关联关系上传 API请求
alibaba.alihealth.drug.kyt.uploadrelation

关联关系上传 */
type AlibabaAlihealthDrugKytUploadrelationAPIRequest struct {
	model.Params
	// 关联关系文件信息
	_saveCodeRelation *SaveCodeRelationType
	// affirmFlag
	_affirmFlag string
	// fileContent
	_fileContent string
	// 加密之后的文件内容字符串
	_fileContentString string
	// 上传文件的企业ID
	_refEntId string
	// 客户端类型
	_clientType string
}

// New
