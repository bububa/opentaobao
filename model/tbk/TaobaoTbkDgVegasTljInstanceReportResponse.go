package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金发放及使用报表 APIResponse
taobao.tbk.dg.vegas.tlj.instance.report

淘礼金实例维度相关报表数据查询
*/
type TaobaoTbkDgVegasTljInstanceReportAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgVegasTljInstanceReportResponse
}

type TaobaoTbkDgVegasTljInstanceReportResponse struct {
    XMLName xml.Name `xml:"tbk_dg_vegas_tlj_instance_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoTbkDgVegasTljInstanceReportResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
