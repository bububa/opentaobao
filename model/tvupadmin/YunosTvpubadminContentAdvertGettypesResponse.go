package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告位类型 API返回值 
yunos.tvpubadmin.content.advert.gettypes

获取广告位类型
*/
type YunosTvpubadminContentAdvertGettypesAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAdvertGettypesResponse
}

// 获取广告位类型 成功返回结果
type YunosTvpubadminContentAdvertGettypesResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_gettypes_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Map<Integer, String> json格式
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
