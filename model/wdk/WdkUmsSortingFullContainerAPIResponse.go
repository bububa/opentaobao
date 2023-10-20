package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsSortingFullContainerAPIResponse dps分货-满箱 API返回值
// wdk.ums.sorting.full.container
//
// dps分货-满箱
type WdkUmsSortingFullContainerAPIResponse struct {
	model.CommonResponse
	WdkUmsSortingFullContainerAPIResponseModel
}

// Reset 清空结构体
func (m *WdkUmsSortingFullContainerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkUmsSortingFullContainerAPIResponseModel).Reset()
}

// WdkUmsSortingFullContainerAPIResponseModel is dps分货-满箱 成功返回结果
type WdkUmsSortingFullContainerAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_sorting_full_container_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkUmsSortingFullContainerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkUmsSortingFullContainerAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkUmsSortingFullContainerAPIResponse)
	},
}

// GetWdkUmsSortingFullContainerAPIResponse 从 sync.Pool 获取 WdkUmsSortingFullContainerAPIResponse
func GetWdkUmsSortingFullContainerAPIResponse() *WdkUmsSortingFullContainerAPIResponse {
	return poolWdkUmsSortingFullContainerAPIResponse.Get().(*WdkUmsSortingFullContainerAPIResponse)
}

// ReleaseWdkUmsSortingFullContainerAPIResponse 将 WdkUmsSortingFullContainerAPIResponse 保存到 sync.Pool
func ReleaseWdkUmsSortingFullContainerAPIResponse(v *WdkUmsSortingFullContainerAPIResponse) {
	v.Reset()
	poolWdkUmsSortingFullContainerAPIResponse.Put(v)
}
