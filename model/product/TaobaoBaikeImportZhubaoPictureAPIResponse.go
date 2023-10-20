package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaikeImportZhubaoPictureAPIResponse 百科图片数据导入 API返回值
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
type TaobaoBaikeImportZhubaoPictureAPIResponse struct {
	model.CommonResponse
	TaobaoBaikeImportZhubaoPictureAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaikeImportZhubaoPictureAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaikeImportZhubaoPictureAPIResponseModel).Reset()
}

// TaobaoBaikeImportZhubaoPictureAPIResponseModel is 百科图片数据导入 成功返回结果
type TaobaoBaikeImportZhubaoPictureAPIResponseModel struct {
	XMLName xml.Name `xml:"baike_import_zhubao_picture_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoBaikeImportZhubaoPictureResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaikeImportZhubaoPictureAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaikeImportZhubaoPictureAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaikeImportZhubaoPictureAPIResponse)
	},
}

// GetTaobaoBaikeImportZhubaoPictureAPIResponse 从 sync.Pool 获取 TaobaoBaikeImportZhubaoPictureAPIResponse
func GetTaobaoBaikeImportZhubaoPictureAPIResponse() *TaobaoBaikeImportZhubaoPictureAPIResponse {
	return poolTaobaoBaikeImportZhubaoPictureAPIResponse.Get().(*TaobaoBaikeImportZhubaoPictureAPIResponse)
}

// ReleaseTaobaoBaikeImportZhubaoPictureAPIResponse 将 TaobaoBaikeImportZhubaoPictureAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaikeImportZhubaoPictureAPIResponse(v *TaobaoBaikeImportZhubaoPictureAPIResponse) {
	v.Reset()
	poolTaobaoBaikeImportZhubaoPictureAPIResponse.Put(v)
}
