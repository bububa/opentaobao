package baichuanctg

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgToutiaoContentAPIResponse 微博输出头条数据 API返回值
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
type AlibabaBaichuanCtgToutiaoContentAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgToutiaoContentAPIResponseModel
}

// AlibabaBaichuanCtgToutiaoContentAPIResponseModel is 微博输出头条数据 成功返回结果
type AlibabaBaichuanCtgToutiaoContentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_toutiao_content_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 内容总体结构
	Result *CtgResponse `json:"result,omitempty" xml:"result,omitempty"`
}
