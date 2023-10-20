package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductquantitygetAPIResponse 渠道库存查询接口 API返回值
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
type TmallsupplychainchannelproductquantitygetAPIResponse struct {
	model.CommonResponse
	TmallsupplychainchannelproductquantitygetAPIResponseModel
}

// TmallsupplychainchannelproductquantitygetAPIResponseModel is 渠道库存查询接口 成功返回结果
type TmallsupplychainchannelproductquantitygetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_quantity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallsupplychainchannelproductquantitygetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
