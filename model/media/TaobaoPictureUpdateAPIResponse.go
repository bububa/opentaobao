package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureUpdateAPIResponse 修改图片名字 API返回值
// taobao.picture.update
//
// 修改指定图片的图片名
type TaobaoPictureUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPictureUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureUpdateAPIResponseModel).Reset()
}

// TaobaoPictureUpdateAPIResponseModel is 修改图片名字 成功返回结果
type TaobaoPictureUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新是否成功
	Done bool `json:"done,omitempty" xml:"done,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Done = false
}

var poolTaobaoPictureUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureUpdateAPIResponse)
	},
}

// GetTaobaoPictureUpdateAPIResponse 从 sync.Pool 获取 TaobaoPictureUpdateAPIResponse
func GetTaobaoPictureUpdateAPIResponse() *TaobaoPictureUpdateAPIResponse {
	return poolTaobaoPictureUpdateAPIResponse.Get().(*TaobaoPictureUpdateAPIResponse)
}

// ReleaseTaobaoPictureUpdateAPIResponse 将 TaobaoPictureUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureUpdateAPIResponse(v *TaobaoPictureUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPictureUpdateAPIResponse.Put(v)
}
