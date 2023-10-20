package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargededutupdateAPIResponse 储值消费 API返回值
// alibaba.alsc.crm.recharge.dedut.update
//
// 增加储值消费接口
type AlibabaalsccrmrechargededutupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmrechargededutupdateAPIResponseModel
}

// AlibabaalsccrmrechargededutupdateAPIResponseModel is 储值消费 成功返回结果
type AlibabaalsccrmrechargededutupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_dedut_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
