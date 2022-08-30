package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemUploadAPIResponse 暗拍发布/编辑B2B商品 API返回值
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
type AlibabaIdleTenderBtobItemUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderBtobItemUploadAPIResponseModel
}

// AlibabaIdleTenderBtobItemUploadAPIResponseModel is 暗拍发布/编辑B2B商品 成功返回结果
type AlibabaIdleTenderBtobItemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
