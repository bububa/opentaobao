package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbcommonaddAPIResponse 通用调用top接口 API返回值
// taobao.xhotel.bnbcommon.add
//
// 通用调用top接口
type TaobaoxhotelbnbcommonaddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbcommonaddAPIResponseModel
}

// TaobaoxhotelbnbcommonaddAPIResponseModel is 通用调用top接口 成功返回结果
type TaobaoxhotelbnbcommonaddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbcommon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	Module string `json:"module,omitempty" xml:"module,omitempty"`
}
