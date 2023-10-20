package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtitemmainsynchronizeAPIResponse 家装新零售主商品同步至阿里 API返回值
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
type TmallnrtitemmainsynchronizeAPIResponse struct {
	model.CommonResponse
	TmallnrtitemmainsynchronizeAPIResponseModel
}

// TmallnrtitemmainsynchronizeAPIResponseModel is 家装新零售主商品同步至阿里 成功返回结果
type TmallnrtitemmainsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_item_main_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtItemMainSynchronize *TmallnrtitemmainsynchronizeResultDo `json:"tmall_nrt_item_main_synchronize,omitempty" xml:"tmall_nrt_item_main_synchronize,omitempty"`
}
