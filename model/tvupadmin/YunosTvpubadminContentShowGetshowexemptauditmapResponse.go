package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松批量查询节目某个牌照的免审状态 API返回值 
yunos.tvpubadmin.content.show.getshowexemptauditmap

迎客松批量查询节目某个牌照的免审状态
*/
type YunosTvpubadminContentShowGetshowexemptauditmapAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentShowGetshowexemptauditmapResponse
}

// 迎客松批量查询节目某个牌照的免审状态 成功返回结果
type YunosTvpubadminContentShowGetshowexemptauditmapResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getshowexemptauditmap_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
