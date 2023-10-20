package dt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAdsDataQueryAPIResponse 导入数据查询 API返回值
// taobao.ads.data.query
//
// 导入数据查询
type TaobaoAdsDataQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAdsDataQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAdsDataQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAdsDataQueryAPIResponseModel).Reset()
}

// TaobaoAdsDataQueryAPIResponseModel is 导入数据查询 成功返回结果
type TaobaoAdsDataQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"ads_data_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 导入数据信息
	RetValue []ExternalTaskDataImportDto `json:"ret_value,omitempty" xml:"ret_value>external_task_data_import_dto,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 0:成功/-1:失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAdsDataQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetValue = m.RetValue[:0]
	m.RetMsg = ""
	m.RetCode = 0
}

var poolTaobaoAdsDataQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAdsDataQueryAPIResponse)
	},
}

// GetTaobaoAdsDataQueryAPIResponse 从 sync.Pool 获取 TaobaoAdsDataQueryAPIResponse
func GetTaobaoAdsDataQueryAPIResponse() *TaobaoAdsDataQueryAPIResponse {
	return poolTaobaoAdsDataQueryAPIResponse.Get().(*TaobaoAdsDataQueryAPIResponse)
}

// ReleaseTaobaoAdsDataQueryAPIResponse 将 TaobaoAdsDataQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAdsDataQueryAPIResponse(v *TaobaoAdsDataQueryAPIResponse) {
	v.Reset()
	poolTaobaoAdsDataQueryAPIResponse.Put(v)
}
