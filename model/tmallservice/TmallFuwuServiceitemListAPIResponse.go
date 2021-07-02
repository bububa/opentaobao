package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuServiceitemListAPIResponse 获取服务商品扩展信息 API返回值
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
type TmallFuwuServiceitemListAPIResponse struct {
	model.CommonResponse
	TmallFuwuServiceitemListAPIResponseModel
}

// TmallFuwuServiceitemListAPIResponseModel is 获取服务商品扩展信息 成功返回结果
type TmallFuwuServiceitemListAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_serviceitem_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallFuwuServiceitemListResult `json:"result,omitempty" xml:"result,omitempty"`
}
