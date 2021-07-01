package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客人PMS账单信息查询 API返回值 
taobao.xhotel.pms.guestbill.get.vtwo

从pms获取客人账单信息
*/
type TaobaoXhotelPmsGuestbillGetVtwoAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel
}

// 客人PMS账单信息查询 成功返回结果
type TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_pms_guestbill_get_vtwo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果集
    Result   *TaobaoXhotelPmsGuestbillGetVtwoResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
