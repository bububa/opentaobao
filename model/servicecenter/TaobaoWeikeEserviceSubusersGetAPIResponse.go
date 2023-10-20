package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceSubusersGetAPIResponse 客服外包订单分配的商家子账号列表 API返回值
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
type TaobaoWeikeEserviceSubusersGetAPIResponse struct {
	model.CommonResponse
	TaobaoWeikeEserviceSubusersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeikeEserviceSubusersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeikeEserviceSubusersGetAPIResponseModel).Reset()
}

// TaobaoWeikeEserviceSubusersGetAPIResponseModel is 客服外包订单分配的商家子账号列表 成功返回结果
type TaobaoWeikeEserviceSubusersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_eservice_subusers_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商家子账号查询结果
	Result *AuthorizedAccountWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeikeEserviceSubusersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWeikeEserviceSubusersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeikeEserviceSubusersGetAPIResponse)
	},
}

// GetTaobaoWeikeEserviceSubusersGetAPIResponse 从 sync.Pool 获取 TaobaoWeikeEserviceSubusersGetAPIResponse
func GetTaobaoWeikeEserviceSubusersGetAPIResponse() *TaobaoWeikeEserviceSubusersGetAPIResponse {
	return poolTaobaoWeikeEserviceSubusersGetAPIResponse.Get().(*TaobaoWeikeEserviceSubusersGetAPIResponse)
}

// ReleaseTaobaoWeikeEserviceSubusersGetAPIResponse 将 TaobaoWeikeEserviceSubusersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeikeEserviceSubusersGetAPIResponse(v *TaobaoWeikeEserviceSubusersGetAPIResponse) {
	v.Reset()
	poolTaobaoWeikeEserviceSubusersGetAPIResponse.Put(v)
}
