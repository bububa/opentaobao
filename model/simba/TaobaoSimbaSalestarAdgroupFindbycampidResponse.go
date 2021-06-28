package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(销量明星)批量获取推广计划下的推广组信息 APIResponse
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组
*/
type TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarAdgroupFindbycampidResponse
}

type TaobaoSimbaSalestarAdgroupFindbycampidResponse struct {
    XMLName xml.Name `xml:"simba_salestar_adgroup_findbycampid_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的推广组分页对象
    
    Adgroups   *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`

    
}
