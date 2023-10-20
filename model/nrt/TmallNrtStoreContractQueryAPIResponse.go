package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstorecontractqueryAPIResponse 摊位合同查询接口 API返回值
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
type TmallnrtstorecontractqueryAPIResponse struct {
	model.CommonResponse
	TmallnrtstorecontractqueryAPIResponseModel
}

// TmallnrtstorecontractqueryAPIResponseModel is 摊位合同查询接口 成功返回结果
type TmallnrtstorecontractqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_contract_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数对象
	Data []EaStoreContractDto `json:"data,omitempty" xml:"data>ea_store_contract_dto,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 成功与否
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
