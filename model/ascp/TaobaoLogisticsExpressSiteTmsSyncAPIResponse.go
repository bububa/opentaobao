package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressSiteTmsSyncAPIResponse 配服务商网点信息同步 API返回值
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
type TaobaoLogisticsExpressSiteTmsSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressSiteTmsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressSiteTmsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressSiteTmsSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressSiteTmsSyncAPIResponseModel is 配服务商网点信息同步 成功返回结果
type TaobaoLogisticsExpressSiteTmsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_site_tms_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	SiteUpsetResponse *SiteUpsetResponse `json:"site_upset_response,omitempty" xml:"site_upset_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressSiteTmsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SiteUpsetResponse = nil
}

var poolTaobaoLogisticsExpressSiteTmsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressSiteTmsSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressSiteTmsSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressSiteTmsSyncAPIResponse
func GetTaobaoLogisticsExpressSiteTmsSyncAPIResponse() *TaobaoLogisticsExpressSiteTmsSyncAPIResponse {
	return poolTaobaoLogisticsExpressSiteTmsSyncAPIResponse.Get().(*TaobaoLogisticsExpressSiteTmsSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressSiteTmsSyncAPIResponse 将 TaobaoLogisticsExpressSiteTmsSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressSiteTmsSyncAPIResponse(v *TaobaoLogisticsExpressSiteTmsSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressSiteTmsSyncAPIResponse.Put(v)
}
