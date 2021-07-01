package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷OTT端广告素材查询 API返回值 
yunos.tvpubadmin.adm.ott.query

查询广告素材
*/
type YunosTvpubadminAdmOttQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminAdmOttQueryAPIResponseModel
}

// 优酷OTT端广告素材查询 成功返回结果
type YunosTvpubadminAdmOttQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_adm_ott_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据详情，json格式
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
