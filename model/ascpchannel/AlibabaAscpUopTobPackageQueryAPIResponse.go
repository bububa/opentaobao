package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
B2B包裹查询接口 API返回值 
alibaba.ascp.uop.tob.package.query

供应链中台TOB包裹查询接口
*/
type AlibabaAscpUopTobPackageQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpUopTobPackageQueryAPIResponseModel
}

// B2B包裹查询接口 成功返回结果
type AlibabaAscpUopTobPackageQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_uop_tob_package_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    PackageQueryResponse   *ResultWrapper `json:"package_query_response,omitempty" xml:"package_query_response,omitempty"`
}
