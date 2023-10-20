package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobankgrouplistAPIResponse 图片银行分组信息获取 API返回值
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
type AlibabaicbuphotobankgrouplistAPIResponse struct {
	model.CommonResponse
	AlibabaicbuphotobankgrouplistAPIResponseModel
}

// AlibabaicbuphotobankgrouplistAPIResponseModel is 图片银行分组信息获取 成功返回结果
type AlibabaicbuphotobankgrouplistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_group_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// groups
	Groups []PhotoAlbumGroup `json:"groups,omitempty" xml:"groups>photo_album_group,omitempty"`
}
