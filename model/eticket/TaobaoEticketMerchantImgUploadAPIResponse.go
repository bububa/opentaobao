package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantImgUploadAPIResponse 码商上传二维码图片 API返回值
// taobao.eticket.merchant.img.upload
//
// 电子凭证的码商可以通过这个接口，上传二维码图片
type TaobaoEticketMerchantImgUploadAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantImgUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantImgUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantImgUploadAPIResponseModel).Reset()
}

// TaobaoEticketMerchantImgUploadAPIResponseModel is 码商上传二维码图片 成功返回结果
type TaobaoEticketMerchantImgUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_img_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 回复对象
	RespBody *UploadImgCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantImgUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantImgUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantImgUploadAPIResponse)
	},
}

// GetTaobaoEticketMerchantImgUploadAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantImgUploadAPIResponse
func GetTaobaoEticketMerchantImgUploadAPIResponse() *TaobaoEticketMerchantImgUploadAPIResponse {
	return poolTaobaoEticketMerchantImgUploadAPIResponse.Get().(*TaobaoEticketMerchantImgUploadAPIResponse)
}

// ReleaseTaobaoEticketMerchantImgUploadAPIResponse 将 TaobaoEticketMerchantImgUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantImgUploadAPIResponse(v *TaobaoEticketMerchantImgUploadAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantImgUploadAPIResponse.Put(v)
}
