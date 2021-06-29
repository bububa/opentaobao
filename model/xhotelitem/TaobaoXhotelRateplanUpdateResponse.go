package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan更新或添加 API返回值 
taobao.xhotel.rateplan.update

酒店产品库rateplan更新或添加
*/
type TaobaoXhotelRateplanUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanUpdateResponse
}

// 价格计划rateplan更新或添加 成功返回结果
type TaobaoXhotelRateplanUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_rateplan_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改的rp id
    Rpid   int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
}
