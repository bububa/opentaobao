package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增产品线 APIResponse
taobao.fenxiao.productcat.add

新增产品线
*/
type TaobaoFenxiaoProductcatAddAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductcatAddResponse
}

type TaobaoFenxiaoProductcatAddResponse struct {
    XMLName xml.Name `xml:"fenxiao_productcat_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 产品线ID
    
    ProductLineId   int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`

    
}
