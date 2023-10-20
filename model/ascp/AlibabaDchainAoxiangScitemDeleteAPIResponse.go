package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemDeleteAPIResponse 货品删除 API返回值
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
type AlibabaDchainAoxiangScitemDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangScitemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangScitemDeleteAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangScitemDeleteAPIResponseModel is 货品删除 成功返回结果
type AlibabaDchainAoxiangScitemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	DeleteScItemResponse *DeleteScItemResponse `json:"delete_sc_item_response,omitempty" xml:"delete_sc_item_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangScitemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeleteScItemResponse = nil
}

var poolAlibabaDchainAoxiangScitemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangScitemDeleteAPIResponse)
	},
}

// GetAlibabaDchainAoxiangScitemDeleteAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangScitemDeleteAPIResponse
func GetAlibabaDchainAoxiangScitemDeleteAPIResponse() *AlibabaDchainAoxiangScitemDeleteAPIResponse {
	return poolAlibabaDchainAoxiangScitemDeleteAPIResponse.Get().(*AlibabaDchainAoxiangScitemDeleteAPIResponse)
}

// ReleaseAlibabaDchainAoxiangScitemDeleteAPIResponse 将 AlibabaDchainAoxiangScitemDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangScitemDeleteAPIResponse(v *AlibabaDchainAoxiangScitemDeleteAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangScitemDeleteAPIResponse.Put(v)
}
