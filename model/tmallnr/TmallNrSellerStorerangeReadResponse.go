package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店服务范围读取 API返回值 
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围
*/
type TmallNrSellerStorerangeReadAPIResponse struct {
    model.CommonResponse
    TmallNrSellerStorerangeReadResponse
}

// 门店服务范围读取 成功返回结果
type TmallNrSellerStorerangeReadResponse struct {
    XMLName xml.Name `xml:"tmall_nr_seller_storerange_read_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    ResultList   *NrResult `json:"result_list,omitempty" xml:"result_list,omitempty"`
}
