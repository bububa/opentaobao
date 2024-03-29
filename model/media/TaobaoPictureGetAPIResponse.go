package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureGetAPIResponse 获取图片信息 API返回值
// taobao.picture.get
//
// 获取图片信息
type TaobaoPictureGetAPIResponse struct {
	model.CommonResponse
	TaobaoPictureGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureGetAPIResponseModel).Reset()
}

// TaobaoPictureGetAPIResponseModel is 获取图片信息 成功返回结果
type TaobaoPictureGetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片信息列表
	Pictures []Picture `json:"pictures,omitempty" xml:"pictures>picture,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Pictures = m.Pictures[:0]
}

var poolTaobaoPictureGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureGetAPIResponse)
	},
}

// GetTaobaoPictureGetAPIResponse 从 sync.Pool 获取 TaobaoPictureGetAPIResponse
func GetTaobaoPictureGetAPIResponse() *TaobaoPictureGetAPIResponse {
	return poolTaobaoPictureGetAPIResponse.Get().(*TaobaoPictureGetAPIResponse)
}

// ReleaseTaobaoPictureGetAPIResponse 将 TaobaoPictureGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureGetAPIResponse(v *TaobaoPictureGetAPIResponse) {
	v.Reset()
	poolTaobaoPictureGetAPIResponse.Put(v)
}
