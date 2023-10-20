package c2m

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpOrganizationGetinviteinfoAPIResponse 淘小铺机构上下级关系 API返回值
// taobao.sebp.organization.getinviteinfo
//
// 机构人员获取机构上下级关系信息
type TaobaoSebpOrganizationGetinviteinfoAPIResponse struct {
	model.CommonResponse
	TaobaoSebpOrganizationGetinviteinfoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSebpOrganizationGetinviteinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSebpOrganizationGetinviteinfoAPIResponseModel).Reset()
}

// TaobaoSebpOrganizationGetinviteinfoAPIResponseModel is 淘小铺机构上下级关系 成功返回结果
type TaobaoSebpOrganizationGetinviteinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"sebp_organization_getinviteinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoSebpOrganizationGetinviteinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSebpOrganizationGetinviteinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSebpOrganizationGetinviteinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSebpOrganizationGetinviteinfoAPIResponse)
	},
}

// GetTaobaoSebpOrganizationGetinviteinfoAPIResponse 从 sync.Pool 获取 TaobaoSebpOrganizationGetinviteinfoAPIResponse
func GetTaobaoSebpOrganizationGetinviteinfoAPIResponse() *TaobaoSebpOrganizationGetinviteinfoAPIResponse {
	return poolTaobaoSebpOrganizationGetinviteinfoAPIResponse.Get().(*TaobaoSebpOrganizationGetinviteinfoAPIResponse)
}

// ReleaseTaobaoSebpOrganizationGetinviteinfoAPIResponse 将 TaobaoSebpOrganizationGetinviteinfoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSebpOrganizationGetinviteinfoAPIResponse(v *TaobaoSebpOrganizationGetinviteinfoAPIResponse) {
	v.Reset()
	poolTaobaoSebpOrganizationGetinviteinfoAPIResponse.Put(v)
}
