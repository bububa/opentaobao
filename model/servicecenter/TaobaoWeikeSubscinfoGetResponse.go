package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
需求订单查询接口 APIResponse
taobao.weike.subscinfo.get

需求订单查询接口
*/
type TaobaoWeikeSubscinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoWeikeSubscinfoGetResponse
}

type TaobaoWeikeSubscinfoGetResponse struct {
    XMLName xml.Name `xml:"weike_subscinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *SubscInfoWrapper `json:"result,omitempty" xml:"result,omitempty"`

    
}
