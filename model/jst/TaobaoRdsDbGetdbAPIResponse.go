package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbGetdbAPIResponse rds获取RDS的DB API返回值
// taobao.rds.db.getdb
//
// rds获取RDS的DB
type TaobaoRdsDbGetdbAPIResponse struct {
	model.CommonResponse
	TaobaoRdsDbGetdbAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdsDbGetdbAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdsDbGetdbAPIResponseModel).Reset()
}

// TaobaoRdsDbGetdbAPIResponseModel is rds获取RDS的DB 成功返回结果
type TaobaoRdsDbGetdbAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_getdb_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdsDbGetdbResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdsDbGetdbAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdsDbGetdbAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdsDbGetdbAPIResponse)
	},
}

// GetTaobaoRdsDbGetdbAPIResponse 从 sync.Pool 获取 TaobaoRdsDbGetdbAPIResponse
func GetTaobaoRdsDbGetdbAPIResponse() *TaobaoRdsDbGetdbAPIResponse {
	return poolTaobaoRdsDbGetdbAPIResponse.Get().(*TaobaoRdsDbGetdbAPIResponse)
}

// ReleaseTaobaoRdsDbGetdbAPIResponse 将 TaobaoRdsDbGetdbAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdsDbGetdbAPIResponse(v *TaobaoRdsDbGetdbAPIResponse) {
	v.Reset()
	poolTaobaoRdsDbGetdbAPIResponse.Put(v)
}
