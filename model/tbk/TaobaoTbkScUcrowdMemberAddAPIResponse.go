package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdMemberAddAPIResponse 淘宝客-服务商-上传人群数据 API返回值
// taobao.tbk.sc.ucrowd.member.add
//
// 服务商使用。支持淘宝客上传人群标签id对应的会员列表，支持全量和增量多种数据更新方式。
type TaobaoTbkScUcrowdMemberAddAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScUcrowdMemberAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdMemberAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScUcrowdMemberAddAPIResponseModel).Reset()
}

// TaobaoTbkScUcrowdMemberAddAPIResponseModel is 淘宝客-服务商-上传人群数据 成功返回结果
type TaobaoTbkScUcrowdMemberAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_member_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdMemberAddRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdMemberAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScUcrowdMemberAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdMemberAddAPIResponse)
	},
}

// GetTaobaoTbkScUcrowdMemberAddAPIResponse 从 sync.Pool 获取 TaobaoTbkScUcrowdMemberAddAPIResponse
func GetTaobaoTbkScUcrowdMemberAddAPIResponse() *TaobaoTbkScUcrowdMemberAddAPIResponse {
	return poolTaobaoTbkScUcrowdMemberAddAPIResponse.Get().(*TaobaoTbkScUcrowdMemberAddAPIResponse)
}

// ReleaseTaobaoTbkScUcrowdMemberAddAPIResponse 将 TaobaoTbkScUcrowdMemberAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScUcrowdMemberAddAPIResponse(v *TaobaoTbkScUcrowdMemberAddAPIResponse) {
	v.Reset()
	poolTaobaoTbkScUcrowdMemberAddAPIResponse.Put(v)
}
