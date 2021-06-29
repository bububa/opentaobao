package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店信息下发 APIResponse
taobao.xhotel.future.softmodify

未来酒店信息下发，包含PMS订单查询和自助入住
*/
type TaobaoXhotelFutureSoftmodifyAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelFutureSoftmodifyResponse
}

type TaobaoXhotelFutureSoftmodifyResponse struct {
    XMLName xml.Name `xml:"xhotel_future_softmodify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求报文示例：https://open.alitrip.com/docs/doc.htm?spm=a21tt.7629140.0.0.Q8jazn&docType=1&articleId=104398中的自助入住请求示例(升级版)和酒店PMS信息查询 响应报文示例：https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153中的自助checkIn订单信息上传和通用PMS结果查询结果
    
    Context   string `json:"context,omitempty" xml:"context,omitempty"`

    
}
