package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外设备获取技能列表 API返回值 
alibaba.ailabs.bots.skils.list

获取ai开放平台技能列表
*/
type AlibabaAilabsBotsSkilsListAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsBotsSkilsListResponse
}

// 对外设备获取技能列表 成功返回结果
type AlibabaAilabsBotsSkilsListResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_bots_skils_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // {        "skillId": 209,        "invocationName": "中文先生",        "name": "测试34",        "serviceProviders": [          {            "icon": "//arplatform.alicdn.com/images/3/1498910818259.png",            "name": "provider1",          }        ],        "botId": 10,        "iconImgUrl": "//arplatform.alicdn.com/images/244/1501764397807.png",        "longDesc": "中文先生是学中文的好帮手。查中文、查成语、听故事样样行。"      }
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
