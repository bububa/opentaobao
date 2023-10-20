package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureDeleteAPIResponse 删除图片空间图片 API返回值
// taobao.picture.delete
//
// 删除图片空间图片
type TaobaoPictureDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPictureDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureDeleteAPIResponseModel).Reset()
}

// TaobaoPictureDeleteAPIResponseModel is 删除图片空间图片 成功返回结果
type TaobaoPictureDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Success = ""
}

var poolTaobaoPictureDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureDeleteAPIResponse)
	},
}

// GetTaobaoPictureDeleteAPIResponse 从 sync.Pool 获取 TaobaoPictureDeleteAPIResponse
func GetTaobaoPictureDeleteAPIResponse() *TaobaoPictureDeleteAPIResponse {
	return poolTaobaoPictureDeleteAPIResponse.Get().(*TaobaoPictureDeleteAPIResponse)
}

// ReleaseTaobaoPictureDeleteAPIResponse 将 TaobaoPictureDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureDeleteAPIResponse(v *TaobaoPictureDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPictureDeleteAPIResponse.Put(v)
}
