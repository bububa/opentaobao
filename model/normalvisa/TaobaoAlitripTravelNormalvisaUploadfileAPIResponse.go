package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaUploadfileAPIResponse 上传电子签证 API返回值
// taobao.alitrip.travel.normalvisa.uploadfile
//
// 上传电子签证
type TaobaoAlitripTravelNormalvisaUploadfileAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelNormalvisaUploadfileAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaUploadfileAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelNormalvisaUploadfileAPIResponseModel).Reset()
}

// TaobaoAlitripTravelNormalvisaUploadfileAPIResponseModel is 上传电子签证 成功返回结果
type TaobaoAlitripTravelNormalvisaUploadfileAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_uploadfile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果：包含results则代表成功保存
	Result *TaobaoAlitripTravelNormalvisaUploadfileResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaUploadfileAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelNormalvisaUploadfileAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaUploadfileAPIResponse)
	},
}

// GetTaobaoAlitripTravelNormalvisaUploadfileAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaUploadfileAPIResponse
func GetTaobaoAlitripTravelNormalvisaUploadfileAPIResponse() *TaobaoAlitripTravelNormalvisaUploadfileAPIResponse {
	return poolTaobaoAlitripTravelNormalvisaUploadfileAPIResponse.Get().(*TaobaoAlitripTravelNormalvisaUploadfileAPIResponse)
}

// ReleaseTaobaoAlitripTravelNormalvisaUploadfileAPIResponse 将 TaobaoAlitripTravelNormalvisaUploadfileAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaUploadfileAPIResponse(v *TaobaoAlitripTravelNormalvisaUploadfileAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaUploadfileAPIResponse.Put(v)
}
