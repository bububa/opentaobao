package tbuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenuidGetbynickAPIResponse 根据买家nick获取买家openuid API返回值
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
type TaobaoUserOpenuidGetbynickAPIResponse struct {
	model.CommonResponse
	TaobaoUserOpenuidGetbynickAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUserOpenuidGetbynickAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserOpenuidGetbynickAPIResponseModel).Reset()
}

// TaobaoUserOpenuidGetbynickAPIResponseModel is 根据买家nick获取买家openuid 成功返回结果
type TaobaoUserOpenuidGetbynickAPIResponseModel struct {
	XMLName xml.Name `xml:"user_openuid_getbynick_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 买家uid对象
	OpenUids []OpenUidInfo `json:"open_uids,omitempty" xml:"open_uids>open_uid_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserOpenuidGetbynickAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUids = m.OpenUids[:0]
}

var poolTaobaoUserOpenuidGetbynickAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserOpenuidGetbynickAPIResponse)
	},
}

// GetTaobaoUserOpenuidGetbynickAPIResponse 从 sync.Pool 获取 TaobaoUserOpenuidGetbynickAPIResponse
func GetTaobaoUserOpenuidGetbynickAPIResponse() *TaobaoUserOpenuidGetbynickAPIResponse {
	return poolTaobaoUserOpenuidGetbynickAPIResponse.Get().(*TaobaoUserOpenuidGetbynickAPIResponse)
}

// ReleaseTaobaoUserOpenuidGetbynickAPIResponse 将 TaobaoUserOpenuidGetbynickAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserOpenuidGetbynickAPIResponse(v *TaobaoUserOpenuidGetbynickAPIResponse) {
	v.Reset()
	poolTaobaoUserOpenuidGetbynickAPIResponse.Put(v)
}
