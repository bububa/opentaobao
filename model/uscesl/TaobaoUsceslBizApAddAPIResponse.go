package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApAddAPIResponse 新增价签通讯AP设备 API返回值
// taobao.uscesl.biz.ap.add
//
// 根据门店和ap的MAC地址新增
type TaobaoUsceslBizApAddAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizApAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizApAddAPIResponseModel).Reset()
}

// TaobaoUsceslBizApAddAPIResponseModel is 新增价签通讯AP设备 成功返回结果
type TaobaoUsceslBizApAddAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoUsceslBizApAddResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizApAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApAddAPIResponse)
	},
}

// GetTaobaoUsceslBizApAddAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizApAddAPIResponse
func GetTaobaoUsceslBizApAddAPIResponse() *TaobaoUsceslBizApAddAPIResponse {
	return poolTaobaoUsceslBizApAddAPIResponse.Get().(*TaobaoUsceslBizApAddAPIResponse)
}

// ReleaseTaobaoUsceslBizApAddAPIResponse 将 TaobaoUsceslBizApAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizApAddAPIResponse(v *TaobaoUsceslBizApAddAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizApAddAPIResponse.Put(v)
}
