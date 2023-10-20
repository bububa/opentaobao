package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizBrandInsertAPIResponse 新增电子价签商家 API返回值
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
type TaobaoUsceslBizBrandInsertAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizBrandInsertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizBrandInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizBrandInsertAPIResponseModel).Reset()
}

// TaobaoUsceslBizBrandInsertAPIResponseModel is 新增电子价签商家 成功返回结果
type TaobaoUsceslBizBrandInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_brand_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizBrandInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslBizBrandInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizBrandInsertAPIResponse)
	},
}

// GetTaobaoUsceslBizBrandInsertAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizBrandInsertAPIResponse
func GetTaobaoUsceslBizBrandInsertAPIResponse() *TaobaoUsceslBizBrandInsertAPIResponse {
	return poolTaobaoUsceslBizBrandInsertAPIResponse.Get().(*TaobaoUsceslBizBrandInsertAPIResponse)
}

// ReleaseTaobaoUsceslBizBrandInsertAPIResponse 将 TaobaoUsceslBizBrandInsertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizBrandInsertAPIResponse(v *TaobaoUsceslBizBrandInsertAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizBrandInsertAPIResponse.Put(v)
}
