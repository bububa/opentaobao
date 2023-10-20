package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbroomtypeAddAPIResponse 民宿新增房源 API返回值
// taobao.xhotel.bnbroomtype.add
//
// 添加民宿房源
type TaobaoXhotelBnbroomtypeAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbroomtypeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbroomtypeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbroomtypeAddAPIResponseModel).Reset()
}

// TaobaoXhotelBnbroomtypeAddAPIResponseModel is 民宿新增房源 成功返回结果
type TaobaoXhotelBnbroomtypeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbroomtype_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房源信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbroomtypeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xroomtype = nil
}

var poolTaobaoXhotelBnbroomtypeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbroomtypeAddAPIResponse)
	},
}

// GetTaobaoXhotelBnbroomtypeAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbroomtypeAddAPIResponse
func GetTaobaoXhotelBnbroomtypeAddAPIResponse() *TaobaoXhotelBnbroomtypeAddAPIResponse {
	return poolTaobaoXhotelBnbroomtypeAddAPIResponse.Get().(*TaobaoXhotelBnbroomtypeAddAPIResponse)
}

// ReleaseTaobaoXhotelBnbroomtypeAddAPIResponse 将 TaobaoXhotelBnbroomtypeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbroomtypeAddAPIResponse(v *TaobaoXhotelBnbroomtypeAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbroomtypeAddAPIResponse.Put(v)
}
