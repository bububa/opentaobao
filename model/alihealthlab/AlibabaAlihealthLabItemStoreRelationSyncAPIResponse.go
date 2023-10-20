package alihealthlab

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthlabitemstorerelationsyncAPIResponse 检验检测业务，isv项目门店关系同步 API返回值
// alibaba.alihealth.lab.item.store.relation.sync
//
// 阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
type AlibabaalihealthlabitemstorerelationsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthlabitemstorerelationsyncAPIResponseModel
}

// AlibabaalihealthlabitemstorerelationsyncAPIResponseModel is 检验检测业务，isv项目门店关系同步 成功返回结果
type AlibabaalihealthlabitemstorerelationsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_lab_item_store_relation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知或处理中
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 可读的错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
