package usergrowth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放询问 API返回值 
taobao.usergrowth.delivery.ask

提供给媒体在曝光广告前调用， 返回是否曝光以及曝光的物料信息
*/
type TaobaoUsergrowthDeliveryAskAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDeliveryAskResponse
}

// 广告投放询问 成功返回结果
type TaobaoUsergrowthDeliveryAskResponse struct {
    XMLName xml.Name `xml:"usergrowth_delivery_ask_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否是目标用户
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 素材和出价详情
    Datas   []TaobaoUsergrowthDeliveryAskData `json:"datas,omitempty" xml:"datas>taobao_usergrowth_delivery_ask_data,omitempty"`
    // 目标用户类型， 1： 拉新；2：促活
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
