package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvSpuSearchAPIResponse spu搜索接口 API返回值
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
type AlibabaIdleIsvSpuSearchAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvSpuSearchAPIResponseModel
}

// AlibabaIdleIsvSpuSearchAPIResponseModel is spu搜索接口 成功返回结果
type AlibabaIdleIsvSpuSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_spu_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleIsvSpuSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
