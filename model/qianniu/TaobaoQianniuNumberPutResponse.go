package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV上传数据接口 APIResponse
taobao.qianniu.number.put

ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。
*/
type TaobaoQianniuNumberPutAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuNumberPutResponse
}

type TaobaoQianniuNumberPutResponse struct {
    XMLName xml.Name `xml:"qianniu_number_put_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否上传成功。返回的是个json串，分别表示每条记录是否成功。
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
