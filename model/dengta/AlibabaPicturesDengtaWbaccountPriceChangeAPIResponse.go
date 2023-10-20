package dengta

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse 微博公众号价格变化通知 API返回值
// alibaba.pictures.dengta.wbaccount.price.change
//
// 微博公众号推广价格变更通知接口
type AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaWbaccountPriceChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPicturesDengtaWbaccountPriceChangeAPIResponseModel).Reset()
}

// AlibabaPicturesDengtaWbaccountPriceChangeAPIResponseModel is 微博公众号价格变化通知 成功返回结果
type AlibabaPicturesDengtaWbaccountPriceChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_wbaccount_price_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaWbaccountPriceChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse)
	},
}

// GetAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse 从 sync.Pool 获取 AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse
func GetAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse() *AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse {
	return poolAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse.Get().(*AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse)
}

// ReleaseAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse 将 AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse(v *AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse) {
	v.Reset()
	poolAlibabaPicturesDengtaWbaccountPriceChangeAPIResponse.Put(v)
}
