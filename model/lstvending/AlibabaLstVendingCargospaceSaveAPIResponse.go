package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingCargospaceSaveAPIResponse 自动售卖机货道数据回流 API返回值
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
type AlibabaLstVendingCargospaceSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingCargospaceSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingCargospaceSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingCargospaceSaveAPIResponseModel).Reset()
}

// AlibabaLstVendingCargospaceSaveAPIResponseModel is 自动售卖机货道数据回流 成功返回结果
type AlibabaLstVendingCargospaceSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_cargospace_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *MultiResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingCargospaceSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendingCargospaceSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingCargospaceSaveAPIResponse)
	},
}

// GetAlibabaLstVendingCargospaceSaveAPIResponse 从 sync.Pool 获取 AlibabaLstVendingCargospaceSaveAPIResponse
func GetAlibabaLstVendingCargospaceSaveAPIResponse() *AlibabaLstVendingCargospaceSaveAPIResponse {
	return poolAlibabaLstVendingCargospaceSaveAPIResponse.Get().(*AlibabaLstVendingCargospaceSaveAPIResponse)
}

// ReleaseAlibabaLstVendingCargospaceSaveAPIResponse 将 AlibabaLstVendingCargospaceSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingCargospaceSaveAPIResponse(v *AlibabaLstVendingCargospaceSaveAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingCargospaceSaveAPIResponse.Put(v)
}
