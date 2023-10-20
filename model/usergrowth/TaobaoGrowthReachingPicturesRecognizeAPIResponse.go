package usergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingPicturesRecognizeAPIResponse 图片识别 API返回值
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
type TaobaoGrowthReachingPicturesRecognizeAPIResponse struct {
	model.CommonResponse
	TaobaoGrowthReachingPicturesRecognizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingPicturesRecognizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGrowthReachingPicturesRecognizeAPIResponseModel).Reset()
}

// TaobaoGrowthReachingPicturesRecognizeAPIResponseModel is 图片识别 成功返回结果
type TaobaoGrowthReachingPicturesRecognizeAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_pictures_recognize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 识别结果list
	RegionList []RegionData `json:"region_list,omitempty" xml:"region_list>region_data,omitempty"`
	// 图片id
	PicId string `json:"pic_id,omitempty" xml:"pic_id,omitempty"`
	// 唤端链接
	DeepLink string `json:"deep_link,omitempty" xml:"deep_link,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingPicturesRecognizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RegionList = m.RegionList[:0]
	m.PicId = ""
	m.DeepLink = ""
}

var poolTaobaoGrowthReachingPicturesRecognizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGrowthReachingPicturesRecognizeAPIResponse)
	},
}

// GetTaobaoGrowthReachingPicturesRecognizeAPIResponse 从 sync.Pool 获取 TaobaoGrowthReachingPicturesRecognizeAPIResponse
func GetTaobaoGrowthReachingPicturesRecognizeAPIResponse() *TaobaoGrowthReachingPicturesRecognizeAPIResponse {
	return poolTaobaoGrowthReachingPicturesRecognizeAPIResponse.Get().(*TaobaoGrowthReachingPicturesRecognizeAPIResponse)
}

// ReleaseTaobaoGrowthReachingPicturesRecognizeAPIResponse 将 TaobaoGrowthReachingPicturesRecognizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGrowthReachingPicturesRecognizeAPIResponse(v *TaobaoGrowthReachingPicturesRecognizeAPIResponse) {
	v.Reset()
	poolTaobaoGrowthReachingPicturesRecognizeAPIResponse.Put(v)
}
