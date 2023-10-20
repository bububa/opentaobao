package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreContractSynAPIResponse 喵零合同同步 API返回值
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
type TmallNrtStoreContractSynAPIResponse struct {
	model.CommonResponse
	TmallNrtStoreContractSynAPIResponseModel
}

// TmallNrtStoreContractSynAPIResponseModel is 喵零合同同步 成功返回结果
type TmallNrtStoreContractSynAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_contract_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 摊位id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 成功与否
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
