package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupDeleteAPIResponse 删除一个推广组 API返回值
// taobao.simba.adgroup.delete
//
// 删除一个推广组
type TaobaoSimbaAdgroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupDeleteAPIResponseModel
}

// TaobaoSimbaAdgroupDeleteAPIResponseModel is 删除一个推广组 成功返回结果
type TaobaoSimbaAdgroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}
