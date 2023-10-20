package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkPosItemprodFindterminalAPIResponse 终端配置支持 API返回值
// taobao.lark.pos.itemprod.findterminal
//
// 终端配置支持,读取如果不存在则创建和远程的连接配置并返回
type TaobaoLarkPosItemprodFindterminalAPIResponse struct {
	model.CommonResponse
	TaobaoLarkPosItemprodFindterminalAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLarkPosItemprodFindterminalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLarkPosItemprodFindterminalAPIResponseModel).Reset()
}

// TaobaoLarkPosItemprodFindterminalAPIResponseModel is 终端配置支持 成功返回结果
type TaobaoLarkPosItemprodFindterminalAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_pos_itemprod_findterminal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 终端配置信息响应
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLarkPosItemprodFindterminalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTaobaoLarkPosItemprodFindterminalAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLarkPosItemprodFindterminalAPIResponse)
	},
}

// GetTaobaoLarkPosItemprodFindterminalAPIResponse 从 sync.Pool 获取 TaobaoLarkPosItemprodFindterminalAPIResponse
func GetTaobaoLarkPosItemprodFindterminalAPIResponse() *TaobaoLarkPosItemprodFindterminalAPIResponse {
	return poolTaobaoLarkPosItemprodFindterminalAPIResponse.Get().(*TaobaoLarkPosItemprodFindterminalAPIResponse)
}

// ReleaseTaobaoLarkPosItemprodFindterminalAPIResponse 将 TaobaoLarkPosItemprodFindterminalAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLarkPosItemprodFindterminalAPIResponse(v *TaobaoLarkPosItemprodFindterminalAPIResponse) {
	v.Reset()
	poolTaobaoLarkPosItemprodFindterminalAPIResponse.Put(v)
}
