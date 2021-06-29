package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新创意相关接口 API返回值 
taobao.simba.salestar.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaSalestarCreativeUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarCreativeUpdateResponse
}

// 销量明星更新创意相关接口 成功返回结果
type TaobaoSimbaSalestarCreativeUpdateResponse struct {
    XMLName xml.Name `xml:"simba_salestar_creative_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 创意修改记录对象
    Creativerecord   *CreativeRecord `json:"creativerecord,omitempty" xml:"creativerecord,omitempty"`
}
