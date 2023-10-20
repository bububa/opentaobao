package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbhouseaddAPIResponse 民宿门店信息添加 API返回值
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
type TaobaoxhotelbnbhouseaddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbhouseaddAPIResponseModel
}

// TaobaoxhotelbnbhouseaddAPIResponseModel is 民宿门店信息添加 成功返回结果
type TaobaoxhotelbnbhouseaddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbhouse_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Results []Xhotel `json:"results,omitempty" xml:"results>xhotel,omitempty"`
}
