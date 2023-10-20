package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvspusearchAPIResponse spu搜索接口 API返回值
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
type AlibabaidleisvspusearchAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvspusearchAPIResponseModel
}

// AlibabaidleisvspusearchAPIResponseModel is spu搜索接口 成功返回结果
type AlibabaidleisvspusearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_spu_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidleisvspusearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
