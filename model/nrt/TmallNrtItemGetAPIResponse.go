package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtitemgetAPIResponse 家装新零售商品信息查询 API返回值
// tmall.nrt.item.get
//
// 查询新零售商品信息
type TmallnrtitemgetAPIResponse struct {
	model.CommonResponse
	TmallnrtitemgetAPIResponseModel
}

// TmallnrtitemgetAPIResponseModel is 家装新零售商品信息查询 成功返回结果
type TmallnrtitemgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtItemGet *TmallnrtitemgetResultDo `json:"tmall_nrt_item_get,omitempty" xml:"tmall_nrt_item_get,omitempty"`
}
