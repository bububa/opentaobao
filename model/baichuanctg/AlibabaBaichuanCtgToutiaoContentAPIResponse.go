package baichuanctg

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgToutiaoContentAPIResponse 微博输出头条数据 API返回值
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
type AlibabaBaichuanCtgToutiaoContentAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgToutiaoContentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgToutiaoContentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanCtgToutiaoContentAPIResponseModel).Reset()
}

// AlibabaBaichuanCtgToutiaoContentAPIResponseModel is 微博输出头条数据 成功返回结果
type AlibabaBaichuanCtgToutiaoContentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_toutiao_content_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 内容总体结构
	Result *CtgResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgToutiaoContentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanCtgToutiaoContentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgToutiaoContentAPIResponse)
	},
}

// GetAlibabaBaichuanCtgToutiaoContentAPIResponse 从 sync.Pool 获取 AlibabaBaichuanCtgToutiaoContentAPIResponse
func GetAlibabaBaichuanCtgToutiaoContentAPIResponse() *AlibabaBaichuanCtgToutiaoContentAPIResponse {
	return poolAlibabaBaichuanCtgToutiaoContentAPIResponse.Get().(*AlibabaBaichuanCtgToutiaoContentAPIResponse)
}

// ReleaseAlibabaBaichuanCtgToutiaoContentAPIResponse 将 AlibabaBaichuanCtgToutiaoContentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanCtgToutiaoContentAPIResponse(v *AlibabaBaichuanCtgToutiaoContentAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanCtgToutiaoContentAPIResponse.Put(v)
}
