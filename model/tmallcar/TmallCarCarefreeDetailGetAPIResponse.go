package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarCarefreeDetailGetAPIResponse 查询业务单信息 API返回值
// tmall.car.carefree.detail.get
//
// 查询业务单信息
type TmallCarCarefreeDetailGetAPIResponse struct {
	model.CommonResponse
	TmallCarCarefreeDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarCarefreeDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarCarefreeDetailGetAPIResponseModel).Reset()
}

// TmallCarCarefreeDetailGetAPIResponseModel is 查询业务单信息 成功返回结果
type TmallCarCarefreeDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_carefree_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据对象
	Data *CarefreeDetailInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarCarefreeDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Extra = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Data = nil
}

var poolTmallCarCarefreeDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarCarefreeDetailGetAPIResponse)
	},
}

// GetTmallCarCarefreeDetailGetAPIResponse 从 sync.Pool 获取 TmallCarCarefreeDetailGetAPIResponse
func GetTmallCarCarefreeDetailGetAPIResponse() *TmallCarCarefreeDetailGetAPIResponse {
	return poolTmallCarCarefreeDetailGetAPIResponse.Get().(*TmallCarCarefreeDetailGetAPIResponse)
}

// ReleaseTmallCarCarefreeDetailGetAPIResponse 将 TmallCarCarefreeDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTmallCarCarefreeDetailGetAPIResponse(v *TmallCarCarefreeDetailGetAPIResponse) {
	v.Reset()
	poolTmallCarCarefreeDetailGetAPIResponse.Put(v)
}
