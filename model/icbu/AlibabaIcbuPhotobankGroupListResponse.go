package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组信息获取 APIResponse
alibaba.icbu.photobank.group.list

图片银行分组信息获取
*/
type AlibabaIcbuPhotobankGroupListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuPhotobankGroupListResponse `json:"alibaba_icbu_photobank_group_list_response,omitempty"` 
    AlibabaIcbuPhotobankGroupListResponse
}

/* model for simplify = false
type AlibabaIcbuPhotobankGroupListResponse struct {

    // groups
    
    Groups  struct {
        PhotoAlbumGroup  []PhotoAlbumGroup `json:"photo_album_group,omitempty"`
    } `json:"groups,omitempty"`
    

}
*/

type AlibabaIcbuPhotobankGroupListResponse struct {

    // groups
    Groups   []PhotoAlbumGroup `json:"groups,omitempty"`

}
