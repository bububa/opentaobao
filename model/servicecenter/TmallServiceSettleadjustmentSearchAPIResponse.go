package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentSearchAPIResponse 服务商15分钟获取数据 API返回值
// tmall.service.settleadjustment.search
//
// 天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
type TmallServiceSettleadjustmentSearchAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettleadjustmentSearchAPIResponseModel).Reset()
}

// TmallServiceSettleadjustmentSearchAPIResponseModel is 服务商15分钟获取数据 成功返回结果
type TmallServiceSettleadjustmentSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettleadjustmentSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettleadjustmentSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettleadjustmentSearchAPIResponse)
	},
}

// GetTmallServiceSettleadjustmentSearchAPIResponse 从 sync.Pool 获取 TmallServiceSettleadjustmentSearchAPIResponse
func GetTmallServiceSettleadjustmentSearchAPIResponse() *TmallServiceSettleadjustmentSearchAPIResponse {
	return poolTmallServiceSettleadjustmentSearchAPIResponse.Get().(*TmallServiceSettleadjustmentSearchAPIResponse)
}

// ReleaseTmallServiceSettleadjustmentSearchAPIResponse 将 TmallServiceSettleadjustmentSearchAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettleadjustmentSearchAPIResponse(v *TmallServiceSettleadjustmentSearchAPIResponse) {
	v.Reset()
	poolTmallServiceSettleadjustmentSearchAPIResponse.Put(v)
}
