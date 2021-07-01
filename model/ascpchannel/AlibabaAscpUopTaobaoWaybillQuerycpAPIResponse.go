package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询电子面单开放的CP列表 API返回值 
alibaba.ascp.uop.taobao.waybill.querycp

查询电子面单开放的CP列表
*/
type AlibabaAscpUopTaobaoWaybillQuerycpAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopTaobaoWaybillQuerycpAPIResponseModel
}

// 查询电子面单开放的CP列表 成功返回结果
type AlibabaAscpUopTaobaoWaybillQuerycpAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_taobao_waybill_querycp_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    QueryCpResponse   *ResultWrapper `json:"query_cp_response,omitempty" xml:"query_cp_response,omitempty"`
}
