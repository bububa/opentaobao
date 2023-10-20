package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundImageUploadAPIResponse OpenMall退款图片上传 API返回值
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
type TaobaoOpenmallRefundImageUploadAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallRefundImageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundImageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallRefundImageUploadAPIResponseModel).Reset()
}

// TaobaoOpenmallRefundImageUploadAPIResponseModel is OpenMall退款图片上传 成功返回结果
type TaobaoOpenmallRefundImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片上传对应Token，用于提交留言接口
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallRefundImageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoOpenmallRefundImageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallRefundImageUploadAPIResponse)
	},
}

// GetTaobaoOpenmallRefundImageUploadAPIResponse 从 sync.Pool 获取 TaobaoOpenmallRefundImageUploadAPIResponse
func GetTaobaoOpenmallRefundImageUploadAPIResponse() *TaobaoOpenmallRefundImageUploadAPIResponse {
	return poolTaobaoOpenmallRefundImageUploadAPIResponse.Get().(*TaobaoOpenmallRefundImageUploadAPIResponse)
}

// ReleaseTaobaoOpenmallRefundImageUploadAPIResponse 将 TaobaoOpenmallRefundImageUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallRefundImageUploadAPIResponse(v *TaobaoOpenmallRefundImageUploadAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallRefundImageUploadAPIResponse.Put(v)
}
