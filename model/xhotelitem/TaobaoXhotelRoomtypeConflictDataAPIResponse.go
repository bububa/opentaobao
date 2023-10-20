package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeConflictDataAPIResponse 商家床型冲突数据接口 API返回值
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
type TaobaoXhotelRoomtypeConflictDataAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomtypeConflictDataAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeConflictDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomtypeConflictDataAPIResponseModel).Reset()
}

// TaobaoXhotelRoomtypeConflictDataAPIResponseModel is 商家床型冲突数据接口 成功返回结果
type TaobaoXhotelRoomtypeConflictDataAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_conflict_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelRoomtypeConflictDataResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeConflictDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelRoomtypeConflictDataAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomtypeConflictDataAPIResponse)
	},
}

// GetTaobaoXhotelRoomtypeConflictDataAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomtypeConflictDataAPIResponse
func GetTaobaoXhotelRoomtypeConflictDataAPIResponse() *TaobaoXhotelRoomtypeConflictDataAPIResponse {
	return poolTaobaoXhotelRoomtypeConflictDataAPIResponse.Get().(*TaobaoXhotelRoomtypeConflictDataAPIResponse)
}

// ReleaseTaobaoXhotelRoomtypeConflictDataAPIResponse 将 TaobaoXhotelRoomtypeConflictDataAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomtypeConflictDataAPIResponse(v *TaobaoXhotelRoomtypeConflictDataAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomtypeConflictDataAPIResponse.Put(v)
}
