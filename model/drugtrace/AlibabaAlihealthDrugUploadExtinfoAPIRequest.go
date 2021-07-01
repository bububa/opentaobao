package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugUploadExtinfoAPIRequest
中药饮片及器械对接 API请求
alibaba.alihealth.drug.upload.extinfo

中药饮片及器械对接 */
type AlibabaAlihealthDrugUploadExtinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugId string
	// 批次
	_batchNo string
	// 扩展信息
	_extInfoDto *ExtInfoDto
}

// New
