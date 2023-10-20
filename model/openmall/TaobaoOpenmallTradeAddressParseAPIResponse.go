package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradeaddressparseAPIResponse openmall服务地址区域码解析 API返回值
// taobao.openmall.trade.address.parse
//
// openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
type TaobaoopenmalltradeaddressparseAPIResponse struct {
	model.CommonResponse
	TaobaoopenmalltradeaddressparseAPIResponseModel
}

// TaobaoopenmalltradeaddressparseAPIResponseModel is openmall服务地址区域码解析 成功返回结果
type TaobaoopenmalltradeaddressparseAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_address_parse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一组地址解析结构，解析正确率与地址完整度相关
	Result *TopParseAddressVo `json:"result,omitempty" xml:"result,omitempty"`
}
