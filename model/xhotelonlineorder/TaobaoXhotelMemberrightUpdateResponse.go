package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店会员权益更新操作 API返回值 
taobao.xhotel.memberright.update

当用户在搜索酒店时，我们需要根据用户是否可享有某项权益来进行相应价格的展示或隐藏，因此我们在酒店搜索时就需要判断用户是否享有某项权益。而由于酒店搜索频率过高，为提高搜索性能并降低第三方接口压力，当用户在搜索酒店时，淘宝会通过读取淘宝本地缓存的用户相关权益信息来进行判断。为提高缓存的准确性，当第三方有用户相关权益有变化时，通过调用淘宝此接口来更新淘宝本地缓存。此接口需要采用Top方式调用。
*/
type TaobaoXhotelMemberrightUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelMemberrightUpdateResponse
}

// 酒店会员权益更新操作 成功返回结果
type TaobaoXhotelMemberrightUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_memberright_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
