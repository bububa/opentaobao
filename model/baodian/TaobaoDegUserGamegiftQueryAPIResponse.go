package baodian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegUserGamegiftQueryAPIResponse 用户数娱游戏礼包查询 API返回值
// taobao.deg.user.gamegift.query
//
// 查询用户数娱礼包列表
type TaobaoDegUserGamegiftQueryAPIResponse struct {
	model.CommonResponse
	TaobaoDegUserGamegiftQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegUserGamegiftQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegUserGamegiftQueryAPIResponseModel).Reset()
}

// TaobaoDegUserGamegiftQueryAPIResponseModel is 用户数娱游戏礼包查询 成功返回结果
type TaobaoDegUserGamegiftQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"deg_user_gamegift_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 礼包信息
	Records []GameGiftRecordDto `json:"records,omitempty" xml:"records>game_gift_record_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegUserGamegiftQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Records = m.Records[:0]
}

var poolTaobaoDegUserGamegiftQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegUserGamegiftQueryAPIResponse)
	},
}

// GetTaobaoDegUserGamegiftQueryAPIResponse 从 sync.Pool 获取 TaobaoDegUserGamegiftQueryAPIResponse
func GetTaobaoDegUserGamegiftQueryAPIResponse() *TaobaoDegUserGamegiftQueryAPIResponse {
	return poolTaobaoDegUserGamegiftQueryAPIResponse.Get().(*TaobaoDegUserGamegiftQueryAPIResponse)
}

// ReleaseTaobaoDegUserGamegiftQueryAPIResponse 将 TaobaoDegUserGamegiftQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegUserGamegiftQueryAPIResponse(v *TaobaoDegUserGamegiftQueryAPIResponse) {
	v.Reset()
	poolTaobaoDegUserGamegiftQueryAPIResponse.Put(v)
}
