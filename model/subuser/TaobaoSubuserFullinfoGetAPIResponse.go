package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserFullinfoGetAPIResponse 获取指定账户子账号的详细信息 API返回值
// taobao.subuser.fullinfo.get
//
// 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
type TaobaoSubuserFullinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubuserFullinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubuserFullinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubuserFullinfoGetAPIResponseModel).Reset()
}

// TaobaoSubuserFullinfoGetAPIResponseModel is 获取指定账户子账号的详细信息 成功返回结果
type TaobaoSubuserFullinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subuser_fullinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
	SubFullinfo *SubUserFullInfo `json:"sub_fullinfo,omitempty" xml:"sub_fullinfo,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubuserFullinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubFullinfo = nil
}

var poolTaobaoSubuserFullinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubuserFullinfoGetAPIResponse)
	},
}

// GetTaobaoSubuserFullinfoGetAPIResponse 从 sync.Pool 获取 TaobaoSubuserFullinfoGetAPIResponse
func GetTaobaoSubuserFullinfoGetAPIResponse() *TaobaoSubuserFullinfoGetAPIResponse {
	return poolTaobaoSubuserFullinfoGetAPIResponse.Get().(*TaobaoSubuserFullinfoGetAPIResponse)
}

// ReleaseTaobaoSubuserFullinfoGetAPIResponse 将 TaobaoSubuserFullinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubuserFullinfoGetAPIResponse(v *TaobaoSubuserFullinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoSubuserFullinfoGetAPIResponse.Put(v)
}
