package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallWorkcardIdentifyAPIResponse 工单核销 API返回值
// tmall.workcard.identify
//
// 工单核销，当工单完成以后，通过调用此接口核销
// 可以按照多维度核销工单，
// 电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:&#39;机器码&#39;}
type TmallWorkcardIdentifyAPIResponse struct {
	model.CommonResponse
	TmallWorkcardIdentifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallWorkcardIdentifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallWorkcardIdentifyAPIResponseModel).Reset()
}

// TmallWorkcardIdentifyAPIResponseModel is 工单核销 成功返回结果
type TmallWorkcardIdentifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_workcard_identify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallWorkcardIdentifyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallWorkcardIdentifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallWorkcardIdentifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallWorkcardIdentifyAPIResponse)
	},
}

// GetTmallWorkcardIdentifyAPIResponse 从 sync.Pool 获取 TmallWorkcardIdentifyAPIResponse
func GetTmallWorkcardIdentifyAPIResponse() *TmallWorkcardIdentifyAPIResponse {
	return poolTmallWorkcardIdentifyAPIResponse.Get().(*TmallWorkcardIdentifyAPIResponse)
}

// ReleaseTmallWorkcardIdentifyAPIResponse 将 TmallWorkcardIdentifyAPIResponse 保存到 sync.Pool
func ReleaseTmallWorkcardIdentifyAPIResponse(v *TmallWorkcardIdentifyAPIResponse) {
	v.Reset()
	poolTmallWorkcardIdentifyAPIResponse.Put(v)
}
