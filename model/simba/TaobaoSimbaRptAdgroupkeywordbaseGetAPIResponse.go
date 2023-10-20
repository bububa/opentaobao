package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptadgroupkeywordbasegetAPIResponse 推广组下的词基础报表数据查询(明细数据不分类型查询) API返回值
// taobao.simba.rpt.adgroupkeywordbase.get
//
// 推广组下的词基础报表数据查询(明细数据不分类型查询)
type TaobaosimbarptadgroupkeywordbasegetAPIResponse struct {
	model.CommonResponse
	TaobaosimbarptadgroupkeywordbasegetAPIResponseModel
}

// TaobaosimbarptadgroupkeywordbasegetAPIResponseModel is 推广组下的词基础报表数据查询(明细数据不分类型查询) 成功返回结果
type TaobaosimbarptadgroupkeywordbasegetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupkeywordbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词基础数据返回结果
	RptAdgroupkeywordBaseList string `json:"rpt_adgroupkeyword_base_list,omitempty" xml:"rpt_adgroupkeyword_base_list,omitempty"`
}
