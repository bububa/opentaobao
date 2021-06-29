package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案信息查询 APIResponse
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。
*/
type TaobaoTbkScPublisherInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkScPublisherInfoGetResponse
}

type TaobaoTbkScPublisherInfoGetResponse struct {
    XMLName xml.Name `xml:"tbk_sc_publisher_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *TaobaoTbkScPublisherInfoGetData `json:"data,omitempty" xml:"data,omitempty"`

    
}
