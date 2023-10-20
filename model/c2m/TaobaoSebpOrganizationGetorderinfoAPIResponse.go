package c2m

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpOrganizationGetorderinfoAPIResponse 淘小铺机构订单信息 API返回值
// taobao.sebp.organization.getorderinfo
//
// 淘小铺合作机构获取机构订单信息，用于机构结算使用
type TaobaoSebpOrganizationGetorderinfoAPIResponse struct {
	model.CommonResponse
	TaobaoSebpOrganizationGetorderinfoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSebpOrganizationGetorderinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSebpOrganizationGetorderinfoAPIResponseModel).Reset()
}

// TaobaoSebpOrganizationGetorderinfoAPIResponseModel is 淘小铺机构订单信息 成功返回结果
type TaobaoSebpOrganizationGetorderinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"sebp_organization_getorderinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoSebpOrganizationGetorderinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSebpOrganizationGetorderinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSebpOrganizationGetorderinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSebpOrganizationGetorderinfoAPIResponse)
	},
}

// GetTaobaoSebpOrganizationGetorderinfoAPIResponse 从 sync.Pool 获取 TaobaoSebpOrganizationGetorderinfoAPIResponse
func GetTaobaoSebpOrganizationGetorderinfoAPIResponse() *TaobaoSebpOrganizationGetorderinfoAPIResponse {
	return poolTaobaoSebpOrganizationGetorderinfoAPIResponse.Get().(*TaobaoSebpOrganizationGetorderinfoAPIResponse)
}

// ReleaseTaobaoSebpOrganizationGetorderinfoAPIResponse 将 TaobaoSebpOrganizationGetorderinfoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSebpOrganizationGetorderinfoAPIResponse(v *TaobaoSebpOrganizationGetorderinfoAPIResponse) {
	v.Reset()
	poolTaobaoSebpOrganizationGetorderinfoAPIResponse.Put(v)
}
