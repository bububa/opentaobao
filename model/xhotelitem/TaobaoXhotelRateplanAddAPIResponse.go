package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateplanAddAPIResponse
酒店产品库rateplan添加 API返回值
taobao.xhotel.rateplan.add

酒店产品库rateplan */
type TaobaoXhotelRateplanAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateplanAddAPIResponseModel
}

// TaobaoXhotelRateplanAddAPIResponseModel is 酒店产品库rateplan添加 成功返回结果
type TaobaoXhotelRateplanAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rateplan_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 生成的rp id
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
}
