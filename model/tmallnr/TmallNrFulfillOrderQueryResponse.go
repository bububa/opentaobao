package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商获取品牌商的单笔订单 APIResponse
tmall.nr.fulfill.order.query

零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
*/
type TmallNrFulfillOrderQueryAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillOrderQueryResponse
}

type TmallNrFulfillOrderQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
