package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据 API返回值 
alibaba.mydata.self.product.get

获取客户产品相关表现数据
*/
type AlibabaMydataSelfProductGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataSelfProductGetResponse
}

// 获取客户产品相关表现数据 成功返回结果
type AlibabaMydataSelfProductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_self_product_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 产品效果查询结果
    Result   *ProductEffect `json:"result,omitempty" xml:"result,omitempty"`
}
