package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组信息获取 APIResponse
alibaba.icbu.photobank.group.list

图片银行分组信息获取
*/
type AlibabaIcbuPhotobankGroupListAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuPhotobankGroupListResponse
}

type AlibabaIcbuPhotobankGroupListResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_photobank_group_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // groups
    
    Groups   []PhotoAlbumGroup `json:"groups,omitempty" xml:"groups>photo_album_group,omitempty"`
    
    
}
