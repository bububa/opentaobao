package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷OTT端广告素材查询 APIResponse
yunos.tvpubadmin.adm.ott.query

查询广告素材
*/
type YunosTvpubadminAdmOttQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminAdmOttQueryResponse
}

type YunosTvpubadminAdmOttQueryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_adm_ott_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据详情，json格式
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
