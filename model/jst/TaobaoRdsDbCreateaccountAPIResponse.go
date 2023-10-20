package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbCreateaccountAPIResponse rds创建数据库账户 API返回值
// taobao.rds.db.createaccount
//
// rds创建数据库账户
type TaobaoRdsDbCreateaccountAPIResponse struct {
	model.CommonResponse
	TaobaoRdsDbCreateaccountAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdsDbCreateaccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdsDbCreateaccountAPIResponseModel).Reset()
}

// TaobaoRdsDbCreateaccountAPIResponseModel is rds创建数据库账户 成功返回结果
type TaobaoRdsDbCreateaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_createaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdsDbCreateaccountResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdsDbCreateaccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdsDbCreateaccountAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdsDbCreateaccountAPIResponse)
	},
}

// GetTaobaoRdsDbCreateaccountAPIResponse 从 sync.Pool 获取 TaobaoRdsDbCreateaccountAPIResponse
func GetTaobaoRdsDbCreateaccountAPIResponse() *TaobaoRdsDbCreateaccountAPIResponse {
	return poolTaobaoRdsDbCreateaccountAPIResponse.Get().(*TaobaoRdsDbCreateaccountAPIResponse)
}

// ReleaseTaobaoRdsDbCreateaccountAPIResponse 将 TaobaoRdsDbCreateaccountAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdsDbCreateaccountAPIResponse(v *TaobaoRdsDbCreateaccountAPIResponse) {
	v.Reset()
	poolTaobaoRdsDbCreateaccountAPIResponse.Put(v)
}
