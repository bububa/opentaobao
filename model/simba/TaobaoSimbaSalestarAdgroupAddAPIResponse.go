package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestaradgroupaddAPIResponse (新)创建一个推广组 API返回值
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
type TaobaosimbasalestaradgroupaddAPIResponse struct {
	model.CommonResponse
	TaobaosimbasalestaradgroupaddAPIResponseModel
}

// TaobaosimbasalestaradgroupaddAPIResponseModel is (新)创建一个推广组 成功返回结果
type TaobaosimbasalestaradgroupaddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增加的推广组
	Adgroup *Adgroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}
