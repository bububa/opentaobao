package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改创意与 API返回值 
taobao.simba.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaCreativeUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCreativeUpdateAPIResponseModel
}

// 修改创意与 成功返回结果
type TaobaoSimbaCreativeUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_creative_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 创意修改记录对象
    Creativerecord   *CreativeRecord `json:"creativerecord,omitempty" xml:"creativerecord,omitempty"`
}
