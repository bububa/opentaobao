package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbDeleteAPIResponse RDS数据库删除 API返回值
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
type TaobaoRdsDbDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoRdsDbDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdsDbDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdsDbDeleteAPIResponseModel).Reset()
}

// TaobaoRdsDbDeleteAPIResponseModel is RDS数据库删除 成功返回结果
type TaobaoRdsDbDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除数据库，返回结果对象
	RdsDbInfo *RdsDbInfo `json:"rds_db_info,omitempty" xml:"rds_db_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdsDbDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RdsDbInfo = nil
}

var poolTaobaoRdsDbDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdsDbDeleteAPIResponse)
	},
}

// GetTaobaoRdsDbDeleteAPIResponse 从 sync.Pool 获取 TaobaoRdsDbDeleteAPIResponse
func GetTaobaoRdsDbDeleteAPIResponse() *TaobaoRdsDbDeleteAPIResponse {
	return poolTaobaoRdsDbDeleteAPIResponse.Get().(*TaobaoRdsDbDeleteAPIResponse)
}

// ReleaseTaobaoRdsDbDeleteAPIResponse 将 TaobaoRdsDbDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdsDbDeleteAPIResponse(v *TaobaoRdsDbDeleteAPIResponse) {
	v.Reset()
	poolTaobaoRdsDbDeleteAPIResponse.Put(v)
}
