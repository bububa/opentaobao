package alihealthlab

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabItemSyncAPIResponse 阿里健康检验检测商品发布 API返回值
// alibaba.alihealth.lab.item.sync
//
// iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
type AlibabaAlihealthLabItemSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthLabItemSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabItemSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthLabItemSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthLabItemSyncAPIResponseModel is 阿里健康检验检测商品发布 成功返回结果
type AlibabaAlihealthLabItemSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_lab_item_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知或参数异常
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 可读的结果码（错误码）
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabItemSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultStatus = ""
	m.ResultCode = ""
}

var poolAlibabaAlihealthLabItemSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthLabItemSyncAPIResponse)
	},
}

// GetAlibabaAlihealthLabItemSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthLabItemSyncAPIResponse
func GetAlibabaAlihealthLabItemSyncAPIResponse() *AlibabaAlihealthLabItemSyncAPIResponse {
	return poolAlibabaAlihealthLabItemSyncAPIResponse.Get().(*AlibabaAlihealthLabItemSyncAPIResponse)
}

// ReleaseAlibabaAlihealthLabItemSyncAPIResponse 将 AlibabaAlihealthLabItemSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthLabItemSyncAPIResponse(v *AlibabaAlihealthLabItemSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthLabItemSyncAPIResponse.Put(v)
}
