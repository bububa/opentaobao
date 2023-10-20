package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeSubscinfoGetAPIResponse 需求订单查询接口 API返回值
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
type TaobaoWeikeSubscinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWeikeSubscinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeikeSubscinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeikeSubscinfoGetAPIResponseModel).Reset()
}

// TaobaoWeikeSubscinfoGetAPIResponseModel is 需求订单查询接口 成功返回结果
type TaobaoWeikeSubscinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_subscinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SubscInfoWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeikeSubscinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWeikeSubscinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeikeSubscinfoGetAPIResponse)
	},
}

// GetTaobaoWeikeSubscinfoGetAPIResponse 从 sync.Pool 获取 TaobaoWeikeSubscinfoGetAPIResponse
func GetTaobaoWeikeSubscinfoGetAPIResponse() *TaobaoWeikeSubscinfoGetAPIResponse {
	return poolTaobaoWeikeSubscinfoGetAPIResponse.Get().(*TaobaoWeikeSubscinfoGetAPIResponse)
}

// ReleaseTaobaoWeikeSubscinfoGetAPIResponse 将 TaobaoWeikeSubscinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeikeSubscinfoGetAPIResponse(v *TaobaoWeikeSubscinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoWeikeSubscinfoGetAPIResponse.Put(v)
}
