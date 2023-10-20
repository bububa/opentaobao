package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemOuteridUpdateAPIResponse 天猫商品/SKU商家编码更新接口 API返回值
// tmall.item.outerid.update
//
// 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallItemOuteridUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemOuteridUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemOuteridUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemOuteridUpdateAPIResponseModel).Reset()
}

// TmallItemOuteridUpdateAPIResponseModel is 天猫商品/SKU商家编码更新接口 成功返回结果
type TmallItemOuteridUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_outerid_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商家编码更新结果
	OuteridUpdateResult string `json:"outerid_update_result,omitempty" xml:"outerid_update_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemOuteridUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OuteridUpdateResult = ""
}

var poolTmallItemOuteridUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemOuteridUpdateAPIResponse)
	},
}

// GetTmallItemOuteridUpdateAPIResponse 从 sync.Pool 获取 TmallItemOuteridUpdateAPIResponse
func GetTmallItemOuteridUpdateAPIResponse() *TmallItemOuteridUpdateAPIResponse {
	return poolTmallItemOuteridUpdateAPIResponse.Get().(*TmallItemOuteridUpdateAPIResponse)
}

// ReleaseTmallItemOuteridUpdateAPIResponse 将 TmallItemOuteridUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemOuteridUpdateAPIResponse(v *TmallItemOuteridUpdateAPIResponse) {
	v.Reset()
	poolTmallItemOuteridUpdateAPIResponse.Put(v)
}
