package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbhousedeleteAPIResponse 民宿门店删除接口 API返回值
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
type TaobaoxhotelbnbhousedeleteAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbhousedeleteAPIResponseModel
}

// TaobaoxhotelbnbhousedeleteAPIResponseModel is 民宿门店删除接口 成功返回结果
type TaobaoxhotelbnbhousedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbhouse_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否出错
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否出错
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}
