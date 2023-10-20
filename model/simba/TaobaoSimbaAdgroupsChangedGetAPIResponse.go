package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupschangedgetAPIResponse 分页获取修改的推广组ID和修改时间 API返回值
// taobao.simba.adgroups.changed.get
//
// 分页获取修改的推广组ID和修改时间
type TaobaosimbaadgroupschangedgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbaadgroupschangedgetAPIResponseModel
}

// TaobaosimbaadgroupschangedgetAPIResponseModel is 分页获取修改的推广组ID和修改时间 成功返回结果
type TaobaosimbaadgroupschangedgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroups_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组分页对象
	Adgroups *AdgroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}
