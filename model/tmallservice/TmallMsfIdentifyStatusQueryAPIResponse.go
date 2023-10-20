package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfIdentifyStatusQueryAPIResponse 喵师傅定案核销状态查询 API返回值
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
type TmallMsfIdentifyStatusQueryAPIResponse struct {
	model.CommonResponse
	TmallMsfIdentifyStatusQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMsfIdentifyStatusQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMsfIdentifyStatusQueryAPIResponseModel).Reset()
}

// TmallMsfIdentifyStatusQueryAPIResponseModel is 喵师傅定案核销状态查询 成功返回结果
type TmallMsfIdentifyStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_identify_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果json串，其中identifyDate位核销日期，signTime为签到时间，identifyStatus位核销状态，1代表已经核销，identifyType为核销类型，取值为：0代表未核销,1代表10位核销码核销,2代表订单号核销,3代表手机号核销,4代表4位核销码核销,5代表通过poi核销
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMsfIdentifyStatusQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTmallMsfIdentifyStatusQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMsfIdentifyStatusQueryAPIResponse)
	},
}

// GetTmallMsfIdentifyStatusQueryAPIResponse 从 sync.Pool 获取 TmallMsfIdentifyStatusQueryAPIResponse
func GetTmallMsfIdentifyStatusQueryAPIResponse() *TmallMsfIdentifyStatusQueryAPIResponse {
	return poolTmallMsfIdentifyStatusQueryAPIResponse.Get().(*TmallMsfIdentifyStatusQueryAPIResponse)
}

// ReleaseTmallMsfIdentifyStatusQueryAPIResponse 将 TmallMsfIdentifyStatusQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallMsfIdentifyStatusQueryAPIResponse(v *TmallMsfIdentifyStatusQueryAPIResponse) {
	v.Reset()
	poolTmallMsfIdentifyStatusQueryAPIResponse.Put(v)
}
