package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfVerifyAPIResponse 喵师傅核销接口 API返回值
// tmall.msf.verify
//
// msf服务核销的top接口
type TmallMsfVerifyAPIResponse struct {
	model.CommonResponse
	TmallMsfVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMsfVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMsfVerifyAPIResponseModel).Reset()
}

// TmallMsfVerifyAPIResponseModel is 喵师傅核销接口 成功返回结果
type TmallMsfVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMsfVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTmallMsfVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMsfVerifyAPIResponse)
	},
}

// GetTmallMsfVerifyAPIResponse 从 sync.Pool 获取 TmallMsfVerifyAPIResponse
func GetTmallMsfVerifyAPIResponse() *TmallMsfVerifyAPIResponse {
	return poolTmallMsfVerifyAPIResponse.Get().(*TmallMsfVerifyAPIResponse)
}

// ReleaseTmallMsfVerifyAPIResponse 将 TmallMsfVerifyAPIResponse 保存到 sync.Pool
func ReleaseTmallMsfVerifyAPIResponse(v *TmallMsfVerifyAPIResponse) {
	v.Reset()
	poolTmallMsfVerifyAPIResponse.Put(v)
}
