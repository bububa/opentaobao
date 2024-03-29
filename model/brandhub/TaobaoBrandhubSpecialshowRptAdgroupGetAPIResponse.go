package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse 品牌号品牌特秀单元报表数据查询 API返回值
// taobao.brandhub.specialshow.rpt.adgroup.get
//
// 获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponseModel).Reset()
}

// TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponseModel is 品牌号品牌特秀单元报表数据查询 成功返回结果
type TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brandhub_specialshow_rpt_adgroup_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	AdgroupRptList []TaobaoBrandhubSpecialshowRptAdgroupGetResult `json:"adgroup_rpt_list,omitempty" xml:"adgroup_rpt_list>taobao_brandhub_specialshow_rpt_adgroup_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AdgroupRptList = m.AdgroupRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse)
	},
}

// GetTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse 从 sync.Pool 获取 TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse
func GetTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse() *TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse {
	return poolTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse.Get().(*TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse)
}

// ReleaseTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse 将 TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse(v *TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse.Put(v)
}
