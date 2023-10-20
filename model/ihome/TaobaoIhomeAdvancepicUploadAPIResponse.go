package ihome

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIhomeAdvancepicUploadAPIResponse ihome图片上传 API返回值
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
type TaobaoIhomeAdvancepicUploadAPIResponse struct {
	model.CommonResponse
	TaobaoIhomeAdvancepicUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoIhomeAdvancepicUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoIhomeAdvancepicUploadAPIResponseModel).Reset()
}

// TaobaoIhomeAdvancepicUploadAPIResponseModel is ihome图片上传 成功返回结果
type TaobaoIhomeAdvancepicUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"ihome_advancepic_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *TaobaoIhomeAdvancepicUploadApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoIhomeAdvancepicUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoIhomeAdvancepicUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoIhomeAdvancepicUploadAPIResponse)
	},
}

// GetTaobaoIhomeAdvancepicUploadAPIResponse 从 sync.Pool 获取 TaobaoIhomeAdvancepicUploadAPIResponse
func GetTaobaoIhomeAdvancepicUploadAPIResponse() *TaobaoIhomeAdvancepicUploadAPIResponse {
	return poolTaobaoIhomeAdvancepicUploadAPIResponse.Get().(*TaobaoIhomeAdvancepicUploadAPIResponse)
}

// ReleaseTaobaoIhomeAdvancepicUploadAPIResponse 将 TaobaoIhomeAdvancepicUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoIhomeAdvancepicUploadAPIResponse(v *TaobaoIhomeAdvancepicUploadAPIResponse) {
	v.Reset()
	poolTaobaoIhomeAdvancepicUploadAPIResponse.Put(v)
}
