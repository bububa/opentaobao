package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupidsdeletedgetAPIResponse 获取删除的推广组ID API返回值
// taobao.simba.adgroupids.deleted.get
//
// 获取删除的推广组ID
type TaobaosimbaadgroupidsdeletedgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbaadgroupidsdeletedgetAPIResponseModel
}

// TaobaosimbaadgroupidsdeletedgetAPIResponseModel is 获取删除的推广组ID 成功返回结果
type TaobaosimbaadgroupidsdeletedgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroupids_deleted_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组ID列表
	DeletedAdgroupIds []int64 `json:"deleted_adgroup_ids,omitempty" xml:"deleted_adgroup_ids>int64,omitempty"`
}
