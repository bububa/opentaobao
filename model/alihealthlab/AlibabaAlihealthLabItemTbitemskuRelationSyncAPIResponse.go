package alihealthlab

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 API返回值
// alibaba.alihealth.lab.item.tbitemsku.relation.sync
//
// 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
type AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponseModel is 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 成功返回结果
type AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_lab_item_tbitemsku_relation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 可读的错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// SUCCESS - 成功，FAIL失败，UNKNOWN - 未知或处理中
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultStatus = ""
	m.ResultMsg = ""
}

var poolAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse)
	},
}

// GetAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse
func GetAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse() *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse {
	return poolAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse.Get().(*AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse)
}

// ReleaseAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse 将 AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse(v *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse.Put(v)
}
