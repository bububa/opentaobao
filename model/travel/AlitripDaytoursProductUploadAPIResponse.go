package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripDaytoursProductUploadAPIResponse 境外一日游/多日游 产品维护接口 API返回值
// alitrip.daytours.product.upload
//
// 境外一日游/多日游 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
type AlitripDaytoursProductUploadAPIResponse struct {
	model.CommonResponse
	AlitripDaytoursProductUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripDaytoursProductUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripDaytoursProductUploadAPIResponseModel).Reset()
}

// AlitripDaytoursProductUploadAPIResponseModel is 境外一日游/多日游 产品维护接口 成功返回结果
type AlitripDaytoursProductUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_daytours_product_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品维护结果
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripDaytoursProductUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripDaytoursProductUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripDaytoursProductUploadAPIResponse)
	},
}

// GetAlitripDaytoursProductUploadAPIResponse 从 sync.Pool 获取 AlitripDaytoursProductUploadAPIResponse
func GetAlitripDaytoursProductUploadAPIResponse() *AlitripDaytoursProductUploadAPIResponse {
	return poolAlitripDaytoursProductUploadAPIResponse.Get().(*AlitripDaytoursProductUploadAPIResponse)
}

// ReleaseAlitripDaytoursProductUploadAPIResponse 将 AlitripDaytoursProductUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripDaytoursProductUploadAPIResponse(v *AlitripDaytoursProductUploadAPIResponse) {
	v.Reset()
	poolAlitripDaytoursProductUploadAPIResponse.Put(v)
}
