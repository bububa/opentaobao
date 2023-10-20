package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersPageAPIResponse 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口） API返回值
// taobao.subusers.page
//
// 分页获取指定账户的子账号简易信息列表
// （新isv接入建议使用taobao.sellercenter.subusers.page接口）
type TaobaoSubusersPageAPIResponse struct {
	model.CommonResponse
	TaobaoSubusersPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubusersPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubusersPageAPIResponseModel).Reset()
}

// TaobaoSubusersPageAPIResponseModel is 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口） 成功返回结果
type TaobaoSubusersPageAPIResponseModel struct {
	XMLName xml.Name `xml:"subusers_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号基本信息
	Subaccounts []SubAccountInfo `json:"subaccounts,omitempty" xml:"subaccounts>sub_account_info,omitempty"`
	// isv本次调用传入的页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 本次调用删选条件下的总子账号数量，除以页大小可得出最大页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// isv本次调用传入的页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubusersPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Subaccounts = m.Subaccounts[:0]
	m.PageSize = 0
	m.TotalCount = 0
	m.PageNum = 0
}

var poolTaobaoSubusersPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubusersPageAPIResponse)
	},
}

// GetTaobaoSubusersPageAPIResponse 从 sync.Pool 获取 TaobaoSubusersPageAPIResponse
func GetTaobaoSubusersPageAPIResponse() *TaobaoSubusersPageAPIResponse {
	return poolTaobaoSubusersPageAPIResponse.Get().(*TaobaoSubusersPageAPIResponse)
}

// ReleaseTaobaoSubusersPageAPIResponse 将 TaobaoSubusersPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubusersPageAPIResponse(v *TaobaoSubusersPageAPIResponse) {
	v.Reset()
	poolTaobaoSubusersPageAPIResponse.Put(v)
}
