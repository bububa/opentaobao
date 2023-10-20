package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsSnInfoQueryAPIResponse 查询单据序列号信息 API返回值
// taobao.wlb.wms.sn.info.query
//
// 查询仓库作业的各类单据记录的Sn信息
type TaobaoWlbWmsSnInfoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsSnInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSnInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsSnInfoQueryAPIResponseModel).Reset()
}

// TaobaoWlbWmsSnInfoQueryAPIResponseModel is 查询单据序列号信息 成功返回结果
type TaobaoWlbWmsSnInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_sn_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *TaobaoWlbWmsSnInfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsSnInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbWmsSnInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsSnInfoQueryAPIResponse)
	},
}

// GetTaobaoWlbWmsSnInfoQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsSnInfoQueryAPIResponse
func GetTaobaoWlbWmsSnInfoQueryAPIResponse() *TaobaoWlbWmsSnInfoQueryAPIResponse {
	return poolTaobaoWlbWmsSnInfoQueryAPIResponse.Get().(*TaobaoWlbWmsSnInfoQueryAPIResponse)
}

// ReleaseTaobaoWlbWmsSnInfoQueryAPIResponse 将 TaobaoWlbWmsSnInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsSnInfoQueryAPIResponse(v *TaobaoWlbWmsSnInfoQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsSnInfoQueryAPIResponse.Put(v)
}
