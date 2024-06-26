package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderItemTagOperateAPIResponse 全渠道商品打标与去标 API返回值
// taobao.omniorder.item.tag.operate
//
// 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderItemTagOperateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderItemTagOperateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderItemTagOperateAPIResponseModel).Reset()
}

// TaobaoOmniorderItemTagOperateAPIResponseModel is 全渠道商品打标与去标 成功返回结果
type TaobaoOmniorderItemTagOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_item_tag_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// code 不为 0时有值，代表异常信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0 正常，否则异常
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderItemTagOperateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Code = ""
}

var poolTaobaoOmniorderItemTagOperateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderItemTagOperateAPIResponse)
	},
}

// GetTaobaoOmniorderItemTagOperateAPIResponse 从 sync.Pool 获取 TaobaoOmniorderItemTagOperateAPIResponse
func GetTaobaoOmniorderItemTagOperateAPIResponse() *TaobaoOmniorderItemTagOperateAPIResponse {
	return poolTaobaoOmniorderItemTagOperateAPIResponse.Get().(*TaobaoOmniorderItemTagOperateAPIResponse)
}

// ReleaseTaobaoOmniorderItemTagOperateAPIResponse 将 TaobaoOmniorderItemTagOperateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderItemTagOperateAPIResponse(v *TaobaoOmniorderItemTagOperateAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderItemTagOperateAPIResponse.Put(v)
}
