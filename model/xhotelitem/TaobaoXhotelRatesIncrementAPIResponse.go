package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelratesincrementAPIResponse 价格推送接口（批量增量） API返回值
// taobao.xhotel.rates.increment
//
// Rate库存&amp;价格增量更新接口，用户仅需要更新Rate中发生变化的库存日历&amp;价格日历即可
type TaobaoxhotelratesincrementAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelratesincrementAPIResponseModel
}

// TaobaoxhotelratesincrementAPIResponseModel is 价格推送接口（批量增量） 成功返回结果
type TaobaoxhotelratesincrementAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rates_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// gid和rpid组合数组&lt;br/&gt;gid_rpid
	GidAndRpids []string `json:"gid_and_rpids,omitempty" xml:"gid_and_rpids>string,omitempty"`
}
