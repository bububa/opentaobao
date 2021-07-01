package wenyuvideo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据优酷人物ID获取人物详情页，包含相关影视和相关人物 API返回值 
youku.wenyuvideo.persion.get

根据优酷人物ID获取人物详情页，包含相关影视和相关人物
*/
type YoukuWenyuvideoPersionGetAPIResponse struct {
    model.CommonResponse
    YoukuWenyuvideoPersionGetAPIResponseModel
}

// 根据优酷人物ID获取人物详情页，包含相关影视和相关人物 成功返回结果
type YoukuWenyuvideoPersionGetAPIResponseModel struct {
    XMLName xml.Name `xml:"youku_wenyuvideo_persion_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *YoukuWenyuvideoPersionGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
