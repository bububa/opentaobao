package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportThreeplResourceGetAPIResponse 3PL直邮获取资源列表 API返回值
// taobao.wlb.import.threepl.resource.get
//
// 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportThreeplResourceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportThreeplResourceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportThreeplResourceGetAPIResponseModel).Reset()
}

// TaobaoWlbImportThreeplResourceGetAPIResponseModel is 3PL直邮获取资源列表 成功返回结果
type TaobaoWlbImportThreeplResourceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_import_threepl_resource_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoWlbImportThreeplResourceGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportThreeplResourceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbImportThreeplResourceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportThreeplResourceGetAPIResponse)
	},
}

// GetTaobaoWlbImportThreeplResourceGetAPIResponse 从 sync.Pool 获取 TaobaoWlbImportThreeplResourceGetAPIResponse
func GetTaobaoWlbImportThreeplResourceGetAPIResponse() *TaobaoWlbImportThreeplResourceGetAPIResponse {
	return poolTaobaoWlbImportThreeplResourceGetAPIResponse.Get().(*TaobaoWlbImportThreeplResourceGetAPIResponse)
}

// ReleaseTaobaoWlbImportThreeplResourceGetAPIResponse 将 TaobaoWlbImportThreeplResourceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportThreeplResourceGetAPIResponse(v *TaobaoWlbImportThreeplResourceGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportThreeplResourceGetAPIResponse.Put(v)
}
