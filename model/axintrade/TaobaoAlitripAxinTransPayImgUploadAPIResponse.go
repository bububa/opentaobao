package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayImgUploadAPIResponse 上传图片到支付宝图片空间接口 API返回值
// taobao.alitrip.axin.trans.pay.img.upload
//
// 阿信供销平台-上传图片到支付宝图片空间接口
type TaobaoAlitripAxinTransPayImgUploadAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransPayImgUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayImgUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransPayImgUploadAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransPayImgUploadAPIResponseModel is 上传图片到支付宝图片空间接口 成功返回结果
type TaobaoAlitripAxinTransPayImgUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_img_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransPayImgUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayImgUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransPayImgUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransPayImgUploadAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransPayImgUploadAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransPayImgUploadAPIResponse
func GetTaobaoAlitripAxinTransPayImgUploadAPIResponse() *TaobaoAlitripAxinTransPayImgUploadAPIResponse {
	return poolTaobaoAlitripAxinTransPayImgUploadAPIResponse.Get().(*TaobaoAlitripAxinTransPayImgUploadAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransPayImgUploadAPIResponse 将 TaobaoAlitripAxinTransPayImgUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayImgUploadAPIResponse(v *TaobaoAlitripAxinTransPayImgUploadAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayImgUploadAPIResponse.Put(v)
}
