package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemDescModulesGetAPIResponse 商品描述模块信息获取 API返回值
// tmall.item.desc.modules.get
//
// 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetAPIResponse struct {
	model.CommonResponse
	TmallItemDescModulesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemDescModulesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemDescModulesGetAPIResponseModel).Reset()
}

// TmallItemDescModulesGetAPIResponseModel is 商品描述模块信息获取 成功返回结果
type TmallItemDescModulesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_desc_modules_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述模块信息
	ModularDescInfo *ModularDescInfo `json:"modular_desc_info,omitempty" xml:"modular_desc_info,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemDescModulesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModularDescInfo = nil
}

var poolTmallItemDescModulesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemDescModulesGetAPIResponse)
	},
}

// GetTmallItemDescModulesGetAPIResponse 从 sync.Pool 获取 TmallItemDescModulesGetAPIResponse
func GetTmallItemDescModulesGetAPIResponse() *TmallItemDescModulesGetAPIResponse {
	return poolTmallItemDescModulesGetAPIResponse.Get().(*TmallItemDescModulesGetAPIResponse)
}

// ReleaseTmallItemDescModulesGetAPIResponse 将 TmallItemDescModulesGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemDescModulesGetAPIResponse(v *TmallItemDescModulesGetAPIResponse) {
	v.Reset()
	poolTmallItemDescModulesGetAPIResponse.Put(v)
}
