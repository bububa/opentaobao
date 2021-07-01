package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控修改 API返回值 
yunos.tvpubadmin.content.advert.manageschedule

广告牌照管控修改
*/
type YunosTvpubadminContentAdvertManagescheduleAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAdvertManagescheduleAPIResponseModel
}

// 广告牌照管控修改 成功返回结果
type YunosTvpubadminContentAdvertManagescheduleAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_manageschedule_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 管理广告排期
    Object   int64 `json:"object,omitempty" xml:"object,omitempty"`
}
