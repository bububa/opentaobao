package dt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNrsItemPricetagRecognizeAPIResponse 价签识别 API返回值
// alibaba.nrs.item.pricetag.recognize
//
// 商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
type AlibabaNrsItemPricetagRecognizeAPIResponse struct {
	model.CommonResponse
	AlibabaNrsItemPricetagRecognizeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNrsItemPricetagRecognizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNrsItemPricetagRecognizeAPIResponseModel).Reset()
}

// AlibabaNrsItemPricetagRecognizeAPIResponseModel is 价签识别 成功返回结果
type AlibabaNrsItemPricetagRecognizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nrs_item_pricetag_recognize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	NrsResult *NrsResult `json:"nrs_result,omitempty" xml:"nrs_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNrsItemPricetagRecognizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.NrsResult = nil
}

var poolAlibabaNrsItemPricetagRecognizeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNrsItemPricetagRecognizeAPIResponse)
	},
}

// GetAlibabaNrsItemPricetagRecognizeAPIResponse 从 sync.Pool 获取 AlibabaNrsItemPricetagRecognizeAPIResponse
func GetAlibabaNrsItemPricetagRecognizeAPIResponse() *AlibabaNrsItemPricetagRecognizeAPIResponse {
	return poolAlibabaNrsItemPricetagRecognizeAPIResponse.Get().(*AlibabaNrsItemPricetagRecognizeAPIResponse)
}

// ReleaseAlibabaNrsItemPricetagRecognizeAPIResponse 将 AlibabaNrsItemPricetagRecognizeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNrsItemPricetagRecognizeAPIResponse(v *AlibabaNrsItemPricetagRecognizeAPIResponse) {
	v.Reset()
	poolAlibabaNrsItemPricetagRecognizeAPIResponse.Put(v)
}
