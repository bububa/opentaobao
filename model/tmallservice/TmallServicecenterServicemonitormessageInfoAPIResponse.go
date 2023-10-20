package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageInfoAPIResponse 服务预警单查询接口 API返回值
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
type TmallServicecenterServicemonitormessageInfoAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicemonitormessageInfoAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicemonitormessageInfoAPIResponseModel).Reset()
}

// TmallServicecenterServicemonitormessageInfoAPIResponseModel is 服务预警单查询接口 成功返回结果
type TmallServicecenterServicemonitormessageInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// biz_id的业务类型， 为1，则bizId为工单id biz_type=2 则表示预警单为运营动态配置的规则产生的预警，对应biz_id为工单id
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 预警级别，1、预警 2、警告 3、严重
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 服务类型
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 处理备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 服务预警单创建时间，格式如：2021-11-06 13:12:11
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 预警规则id，对应一种业务触发的条件
	RuleId string `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// 预警提示内容，如请及时处理即将超出6小时时未回传工人信息的服务工单，详情如下：父订单编号：XXX，服务子订单：XXX，服务工单号：XXX
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 预警状态 0、已生成 1、已预警 2、已收到 3、已读 4、处理中 5、已处理
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 业务实体id，参考 biz_type
	BizId int64 `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 预警单主键id ，反馈处理进度和备注时需要回传此ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizType = ""
	m.Level = ""
	m.ServiceCode = ""
	m.Memo = ""
	m.GmtCreate = ""
	m.RuleId = ""
	m.Content = ""
	m.Status = ""
	m.BizId = 0
	m.Id = 0
}

var poolTmallServicecenterServicemonitormessageInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicemonitormessageInfoAPIResponse)
	},
}

// GetTmallServicecenterServicemonitormessageInfoAPIResponse 从 sync.Pool 获取 TmallServicecenterServicemonitormessageInfoAPIResponse
func GetTmallServicecenterServicemonitormessageInfoAPIResponse() *TmallServicecenterServicemonitormessageInfoAPIResponse {
	return poolTmallServicecenterServicemonitormessageInfoAPIResponse.Get().(*TmallServicecenterServicemonitormessageInfoAPIResponse)
}

// ReleaseTmallServicecenterServicemonitormessageInfoAPIResponse 将 TmallServicecenterServicemonitormessageInfoAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageInfoAPIResponse(v *TmallServicecenterServicemonitormessageInfoAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageInfoAPIResponse.Put(v)
}
