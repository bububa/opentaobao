package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScRelationRecordAPIResponse 淘宝客-服务商-私域会员标记 API返回值
// taobao.tbk.sc.relation.record
//
// 服务商使用。支持淘宝客通过入参私域外部ID，获得待私域会员可标记的链接，会员打开该链接后，可帮助媒体自动生成会员运营id进行标记，同时自动跳转到推广落地页。
type TaobaoTbkScRelationRecordAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScRelationRecordAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScRelationRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScRelationRecordAPIResponseModel).Reset()
}

// TaobaoTbkScRelationRecordAPIResponseModel is 淘宝客-服务商-私域会员标记 成功返回结果
type TaobaoTbkScRelationRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_relation_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScRelationRecordMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScRelationRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScRelationRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScRelationRecordAPIResponse)
	},
}

// GetTaobaoTbkScRelationRecordAPIResponse 从 sync.Pool 获取 TaobaoTbkScRelationRecordAPIResponse
func GetTaobaoTbkScRelationRecordAPIResponse() *TaobaoTbkScRelationRecordAPIResponse {
	return poolTaobaoTbkScRelationRecordAPIResponse.Get().(*TaobaoTbkScRelationRecordAPIResponse)
}

// ReleaseTaobaoTbkScRelationRecordAPIResponse 将 TaobaoTbkScRelationRecordAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScRelationRecordAPIResponse(v *TaobaoTbkScRelationRecordAPIResponse) {
	v.Reset()
	poolTaobaoTbkScRelationRecordAPIResponse.Put(v)
}
