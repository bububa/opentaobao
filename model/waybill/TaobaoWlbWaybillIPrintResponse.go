package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
打印确认接口v1.0 APIResponse
taobao.wlb.waybill.i.print

打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
*/
type TaobaoWlbWaybillIPrintAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWaybillIPrintResponse
}

type TaobaoWlbWaybillIPrintResponse struct {
    XMLName xml.Name `xml:"wlb_waybill_i_print_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 面单打印信息
    
    WaybillApplyPrintCheckInfos   []WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos,omitempty" xml:"waybill_apply_print_check_infos>waybill_apply_print_check_info,omitempty"`
    
    
}
