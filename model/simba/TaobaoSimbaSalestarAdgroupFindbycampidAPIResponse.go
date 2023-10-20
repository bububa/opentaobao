package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestaradgroupfindbycampidAPIResponse (销量明星)批量获取推广计划下的推广组信息 API返回值
// taobao.simba.salestar.adgroup.findbycampid
//
// 批量得到推广计划下的推广组
type TaobaosimbasalestaradgroupfindbycampidAPIResponse struct {
	model.CommonResponse
	TaobaosimbasalestaradgroupfindbycampidAPIResponseModel
}

// TaobaosimbasalestaradgroupfindbycampidAPIResponseModel is (销量明星)批量获取推广计划下的推广组信息 成功返回结果
type TaobaosimbasalestaradgroupfindbycampidAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_findbycampid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Adgroups *AdgroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}
