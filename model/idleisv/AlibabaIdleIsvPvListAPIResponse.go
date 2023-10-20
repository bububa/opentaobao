package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvpvlistAPIResponse 闲鱼已验货pv查询 API返回值
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
type AlibabaidleisvpvlistAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvpvlistAPIResponseModel
}

// AlibabaidleisvpvlistAPIResponseModel is 闲鱼已验货pv查询 成功返回结果
type AlibabaidleisvpvlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_pv_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidleisvpvlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
