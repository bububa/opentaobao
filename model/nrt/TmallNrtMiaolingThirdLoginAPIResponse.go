package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMiaolingThirdLoginAPIResponse 喵零第三方免登 API返回值
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
type TmallNrtMiaolingThirdLoginAPIResponse struct {
	model.CommonResponse
	TmallNrtMiaolingThirdLoginAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtMiaolingThirdLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtMiaolingThirdLoginAPIResponseModel).Reset()
}

// TmallNrtMiaolingThirdLoginAPIResponseModel is 喵零第三方免登 成功返回结果
type TmallNrtMiaolingThirdLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_miaoling_third_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtMiaolingThirdLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtMiaolingThirdLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtMiaolingThirdLoginAPIResponse)
	},
}

// GetTmallNrtMiaolingThirdLoginAPIResponse 从 sync.Pool 获取 TmallNrtMiaolingThirdLoginAPIResponse
func GetTmallNrtMiaolingThirdLoginAPIResponse() *TmallNrtMiaolingThirdLoginAPIResponse {
	return poolTmallNrtMiaolingThirdLoginAPIResponse.Get().(*TmallNrtMiaolingThirdLoginAPIResponse)
}

// ReleaseTmallNrtMiaolingThirdLoginAPIResponse 将 TmallNrtMiaolingThirdLoginAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtMiaolingThirdLoginAPIResponse(v *TmallNrtMiaolingThirdLoginAPIResponse) {
	v.Reset()
	poolTmallNrtMiaolingThirdLoginAPIResponse.Put(v)
}
