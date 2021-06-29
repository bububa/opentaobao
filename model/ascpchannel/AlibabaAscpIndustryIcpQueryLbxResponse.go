package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
icp订单号查询lbx订单号 APIResponse
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号
*/
type AlibabaAscpIndustryIcpQueryLbxAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryIcpQueryLbxResponse
}

type AlibabaAscpIndustryIcpQueryLbxResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_icp_query_lbx_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    BizResponse   *ResultWrapper `json:"biz_response,omitempty" xml:"biz_response,omitempty"`

    
}
