package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyskuQueryStatusAPIResponse 查询商家是否开通自助修改商品信息服务 API返回值
// taobao.modifysku.query.status
//
// 查询商家是否开通自助修改商品信息服务
// 1. 没有授权
// 2. 已与当前appkey签约
// 3. 没有签约
// 4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
type TaobaoModifyskuQueryStatusAPIResponse struct {
	model.CommonResponse
	TaobaoModifyskuQueryStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoModifyskuQueryStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoModifyskuQueryStatusAPIResponseModel).Reset()
}

// TaobaoModifyskuQueryStatusAPIResponseModel is 查询商家是否开通自助修改商品信息服务 成功返回结果
type TaobaoModifyskuQueryStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"modifysku_query_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Model *CheckSignSkuResponse `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoModifyskuQueryStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolTaobaoModifyskuQueryStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoModifyskuQueryStatusAPIResponse)
	},
}

// GetTaobaoModifyskuQueryStatusAPIResponse 从 sync.Pool 获取 TaobaoModifyskuQueryStatusAPIResponse
func GetTaobaoModifyskuQueryStatusAPIResponse() *TaobaoModifyskuQueryStatusAPIResponse {
	return poolTaobaoModifyskuQueryStatusAPIResponse.Get().(*TaobaoModifyskuQueryStatusAPIResponse)
}

// ReleaseTaobaoModifyskuQueryStatusAPIResponse 将 TaobaoModifyskuQueryStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoModifyskuQueryStatusAPIResponse(v *TaobaoModifyskuQueryStatusAPIResponse) {
	v.Reset()
	poolTaobaoModifyskuQueryStatusAPIResponse.Put(v)
}
