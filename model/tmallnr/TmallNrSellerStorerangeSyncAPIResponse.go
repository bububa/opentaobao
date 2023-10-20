package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrsellerstorerangesyncAPIResponse 同步商户中心服务范围 API返回值
// tmall.nr.seller.storerange.sync
//
// 同步商户中心服务范围
type TmallnrsellerstorerangesyncAPIResponse struct {
	model.CommonResponse
	TmallnrsellerstorerangesyncAPIResponseModel
}

// TmallnrsellerstorerangesyncAPIResponseModel is 同步商户中心服务范围 成功返回结果
type TmallnrsellerstorerangesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_seller_storerange_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 请求是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}
