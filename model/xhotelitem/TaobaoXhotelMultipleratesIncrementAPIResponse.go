package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelMultipleratesIncrementAPIResponse
复杂房价推送接口（批量增量） API返回值
taobao.xhotel.multiplerates.increment

复杂房价批量增量更新，只会更新指定日期的信息
完全涵盖了taobao.xhotel.rates.increment接口的功能 */
type TaobaoXhotelMultipleratesIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultipleratesIncrementAPIResponseModel
}

// TaobaoXhotelMultipleratesIncrementAPIResponseModel is 复杂房价推送接口（批量增量） 成功返回结果
type TaobaoXhotelMultipleratesIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerates_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品id-房价id-入住人数-连住天数  的集合
	GidAndRpidOccupancyLengthofstay []string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" xml:"gid_and_rpid_occupancy_lengthofstay>string,omitempty"`
	// 批量更新的时候，如果部分更新失败，会展示部分失败的原因
	Warnmessage string `json:"warnmessage,omitempty" xml:"warnmessage,omitempty"`
}
