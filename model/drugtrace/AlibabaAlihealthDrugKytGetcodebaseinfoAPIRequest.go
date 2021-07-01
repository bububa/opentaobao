package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest
码的药品信息查询 API请求
alibaba.alihealth.drug.kyt.getcodebaseinfo

提供根据码查询码基本信息接口 */
type AlibabaAlihealthDrugKytGetcodebaseinfoAPIRequest struct {
	model.Params
	// 码
	_code string
	// 企业唯一标识
	_refEntId string
}

// New
