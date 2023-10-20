package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLocationRelationEditAPIResponse 地点关联关系增量编辑 API返回值
// taobao.location.relation.edit
//
// 地点关联关系增量编辑
type TaobaoLocationRelationEditAPIResponse struct {
	model.CommonResponse
	TaobaoLocationRelationEditAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLocationRelationEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLocationRelationEditAPIResponseModel).Reset()
}

// TaobaoLocationRelationEditAPIResponseModel is 地点关联关系增量编辑 成功返回结果
type TaobaoLocationRelationEditAPIResponseModel struct {
	XMLName xml.Name `xml:"location_relation_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLocationRelationEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Errorcode = ""
	m.IsSuccess = false
}

var poolTaobaoLocationRelationEditAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLocationRelationEditAPIResponse)
	},
}

// GetTaobaoLocationRelationEditAPIResponse 从 sync.Pool 获取 TaobaoLocationRelationEditAPIResponse
func GetTaobaoLocationRelationEditAPIResponse() *TaobaoLocationRelationEditAPIResponse {
	return poolTaobaoLocationRelationEditAPIResponse.Get().(*TaobaoLocationRelationEditAPIResponse)
}

// ReleaseTaobaoLocationRelationEditAPIResponse 将 TaobaoLocationRelationEditAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLocationRelationEditAPIResponse(v *TaobaoLocationRelationEditAPIResponse) {
	v.Reset()
	poolTaobaoLocationRelationEditAPIResponse.Put(v)
}
