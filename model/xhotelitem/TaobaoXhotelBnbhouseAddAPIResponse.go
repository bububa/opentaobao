package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBnbhouseAddAPIResponse
民宿门店信息添加 API返回值
taobao.xhotel.bnbhouse.add

添加和更新民宿门店的信息 */
type TaobaoXhotelBnbhouseAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbhouseAddAPIResponseModel
}

// TaobaoXhotelBnbhouseAddAPIResponseModel is 民宿门店信息添加 成功返回结果
type TaobaoXhotelBnbhouseAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbhouse_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Results []XHotel `json:"results,omitempty" xml:"results>x_hotel,omitempty"`
}
