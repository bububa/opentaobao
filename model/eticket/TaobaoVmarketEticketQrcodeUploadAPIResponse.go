package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketQrcodeUploadAPIResponse 码商二维码图片上传 API返回值
// taobao.vmarket.eticket.qrcode.upload
//
// 电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
type TaobaoVmarketEticketQrcodeUploadAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketQrcodeUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketQrcodeUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketQrcodeUploadAPIResponseModel).Reset()
}

// TaobaoVmarketEticketQrcodeUploadAPIResponseModel is 码商二维码图片上传 成功返回结果
type TaobaoVmarketEticketQrcodeUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_qrcode_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片文件名称
	ImgFilename string `json:"img_filename,omitempty" xml:"img_filename,omitempty"`
	// 1:成功  其它为失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketQrcodeUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ImgFilename = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketQrcodeUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketQrcodeUploadAPIResponse)
	},
}

// GetTaobaoVmarketEticketQrcodeUploadAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketQrcodeUploadAPIResponse
func GetTaobaoVmarketEticketQrcodeUploadAPIResponse() *TaobaoVmarketEticketQrcodeUploadAPIResponse {
	return poolTaobaoVmarketEticketQrcodeUploadAPIResponse.Get().(*TaobaoVmarketEticketQrcodeUploadAPIResponse)
}

// ReleaseTaobaoVmarketEticketQrcodeUploadAPIResponse 将 TaobaoVmarketEticketQrcodeUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketQrcodeUploadAPIResponse(v *TaobaoVmarketEticketQrcodeUploadAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketQrcodeUploadAPIResponse.Put(v)
}
