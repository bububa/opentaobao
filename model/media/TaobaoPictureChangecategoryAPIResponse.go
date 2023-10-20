package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureChangecategoryAPIResponse 修改图片的分类 API返回值
// taobao.picture.changecategory
//
// 把批量的图片移动到某个分类下
type TaobaoPictureChangecategoryAPIResponse struct {
	model.CommonResponse
	TaobaoPictureChangecategoryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureChangecategoryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureChangecategoryAPIResponseModel).Reset()
}

// TaobaoPictureChangecategoryAPIResponseModel is 修改图片的分类 成功返回结果
type TaobaoPictureChangecategoryAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_changecategory_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移动图片是否成功：部分移动成功为true，全部移动失败为false。
	Done bool `json:"done,omitempty" xml:"done,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureChangecategoryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Done = false
}

var poolTaobaoPictureChangecategoryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureChangecategoryAPIResponse)
	},
}

// GetTaobaoPictureChangecategoryAPIResponse 从 sync.Pool 获取 TaobaoPictureChangecategoryAPIResponse
func GetTaobaoPictureChangecategoryAPIResponse() *TaobaoPictureChangecategoryAPIResponse {
	return poolTaobaoPictureChangecategoryAPIResponse.Get().(*TaobaoPictureChangecategoryAPIResponse)
}

// ReleaseTaobaoPictureChangecategoryAPIResponse 将 TaobaoPictureChangecategoryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureChangecategoryAPIResponse(v *TaobaoPictureChangecategoryAPIResponse) {
	v.Reset()
	poolTaobaoPictureChangecategoryAPIResponse.Put(v)
}
