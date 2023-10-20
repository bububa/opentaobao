package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslInfoGetAPIResponse 价签设备信息查询接口 API返回值
// taobao.uscesl.biz.esl.info.get
//
// 价签设备信息查询接口
type TaobaoUsceslBizEslInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizEslInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizEslInfoGetAPIResponseModel).Reset()
}

// TaobaoUsceslBizEslInfoGetAPIResponseModel is 价签设备信息查询接口 成功返回结果
type TaobaoUsceslBizEslInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_esl_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result.sucess表示是否成功，target是设备信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslBizEslInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizEslInfoGetAPIResponse)
	},
}

// GetTaobaoUsceslBizEslInfoGetAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizEslInfoGetAPIResponse
func GetTaobaoUsceslBizEslInfoGetAPIResponse() *TaobaoUsceslBizEslInfoGetAPIResponse {
	return poolTaobaoUsceslBizEslInfoGetAPIResponse.Get().(*TaobaoUsceslBizEslInfoGetAPIResponse)
}

// ReleaseTaobaoUsceslBizEslInfoGetAPIResponse 将 TaobaoUsceslBizEslInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizEslInfoGetAPIResponse(v *TaobaoUsceslBizEslInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizEslInfoGetAPIResponse.Put(v)
}
