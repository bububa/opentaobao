package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubusersGetAPIResponse 查询指定账户的子账号列表 API返回值
// taobao.sellercenter.subusers.get
//
// 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetAPIResponse struct {
	model.CommonResponse
	TaobaoSellercenterSubusersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSellercenterSubusersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSellercenterSubusersGetAPIResponseModel).Reset()
}

// TaobaoSellercenterSubusersGetAPIResponseModel is 查询指定账户的子账号列表 成功返回结果
type TaobaoSellercenterSubusersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_subusers_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
	Subusers []SubUserInfo `json:"subusers,omitempty" xml:"subusers>sub_user_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSellercenterSubusersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Subusers = m.Subusers[:0]
}

var poolTaobaoSellercenterSubusersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSellercenterSubusersGetAPIResponse)
	},
}

// GetTaobaoSellercenterSubusersGetAPIResponse 从 sync.Pool 获取 TaobaoSellercenterSubusersGetAPIResponse
func GetTaobaoSellercenterSubusersGetAPIResponse() *TaobaoSellercenterSubusersGetAPIResponse {
	return poolTaobaoSellercenterSubusersGetAPIResponse.Get().(*TaobaoSellercenterSubusersGetAPIResponse)
}

// ReleaseTaobaoSellercenterSubusersGetAPIResponse 将 TaobaoSellercenterSubusersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSellercenterSubusersGetAPIResponse(v *TaobaoSellercenterSubusersGetAPIResponse) {
	v.Reset()
	poolTaobaoSellercenterSubusersGetAPIResponse.Put(v)
}
