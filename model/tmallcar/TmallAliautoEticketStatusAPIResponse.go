package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStatusAPIResponse 查询电子凭证状态 API返回值
// tmall.aliauto.eticket.status
//
// 查询天猫汽车二轮车行业门店电子凭证状态
type TmallAliautoEticketStatusAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoEticketStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoEticketStatusAPIResponseModel).Reset()
}

// TmallAliautoEticketStatusAPIResponseModel is 查询电子凭证状态 成功返回结果
type TmallAliautoEticketStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *EticketInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoEticketStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTmallAliautoEticketStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoEticketStatusAPIResponse)
	},
}

// GetTmallAliautoEticketStatusAPIResponse 从 sync.Pool 获取 TmallAliautoEticketStatusAPIResponse
func GetTmallAliautoEticketStatusAPIResponse() *TmallAliautoEticketStatusAPIResponse {
	return poolTmallAliautoEticketStatusAPIResponse.Get().(*TmallAliautoEticketStatusAPIResponse)
}

// ReleaseTmallAliautoEticketStatusAPIResponse 将 TmallAliautoEticketStatusAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoEticketStatusAPIResponse(v *TmallAliautoEticketStatusAPIResponse) {
	v.Reset()
	poolTmallAliautoEticketStatusAPIResponse.Put(v)
}
