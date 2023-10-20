package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureIsreferencedGetAPIResponse 图片是否被引用 API返回值
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetAPIResponse struct {
	model.CommonResponse
	TaobaoPictureIsreferencedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureIsreferencedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureIsreferencedGetAPIResponseModel).Reset()
}

// TaobaoPictureIsreferencedGetAPIResponseModel is 图片是否被引用 成功返回结果
type TaobaoPictureIsreferencedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_isreferenced_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片是否被引用
	IsReferenced bool `json:"is_referenced,omitempty" xml:"is_referenced,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureIsreferencedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsReferenced = false
}

var poolTaobaoPictureIsreferencedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureIsreferencedGetAPIResponse)
	},
}

// GetTaobaoPictureIsreferencedGetAPIResponse 从 sync.Pool 获取 TaobaoPictureIsreferencedGetAPIResponse
func GetTaobaoPictureIsreferencedGetAPIResponse() *TaobaoPictureIsreferencedGetAPIResponse {
	return poolTaobaoPictureIsreferencedGetAPIResponse.Get().(*TaobaoPictureIsreferencedGetAPIResponse)
}

// ReleaseTaobaoPictureIsreferencedGetAPIResponse 将 TaobaoPictureIsreferencedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureIsreferencedGetAPIResponse(v *TaobaoPictureIsreferencedGetAPIResponse) {
	v.Reset()
	poolTaobaoPictureIsreferencedGetAPIResponse.Put(v)
}
