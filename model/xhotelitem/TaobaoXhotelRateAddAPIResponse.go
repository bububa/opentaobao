package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelrateaddAPIResponse 新增专享房价 API返回值
// taobao.xhotel.rate.add
//
// 酒店产品库rate添加
type TaobaoxhotelrateaddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelrateaddAPIResponseModel
}

// TaobaoxhotelrateaddAPIResponseModel is 新增专享房价 成功返回结果
type TaobaoxhotelrateaddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// 酒店商品id-酒店rpID
	GidAndRpid string `json:"gid_and_rpid,omitempty" xml:"gid_and_rpid,omitempty"`
	// warnMessage
	WarnMessage string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
}
