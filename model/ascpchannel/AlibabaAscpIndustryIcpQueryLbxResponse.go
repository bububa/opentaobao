package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
icp订单号查询lbx订单号 API返回值 
alibaba.ascp.industry.icp.query.lbx

根据icp订单号查询lbx订单号
*/
type AlibabaAscpIndustryIcpQueryLbxAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryIcpQueryLbxResponse
}

// icp订单号查询lbx订单号 成功返回结果
type AlibabaAscpIndustryIcpQueryLbxResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_icp_query_lbx_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    BizResponse   *ResultWrapper `json:"biz_response,omitempty" xml:"biz_response,omitempty"`
}
