package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoLoginUserGetAPIResponse 获取分销用户登录信息 API返回值
// taobao.fenxiao.login.user.get
//
// 获取用户登录信息
type TaobaoFenxiaoLoginUserGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoLoginUserGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoLoginUserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoLoginUserGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoLoginUserGetAPIResponseModel is 获取分销用户登录信息 成功返回结果
type TaobaoFenxiaoLoginUserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_login_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 登录用户信息
	LoginUser *LoginUser `json:"login_user,omitempty" xml:"login_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoLoginUserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.LoginUser = nil
}

var poolTaobaoFenxiaoLoginUserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoLoginUserGetAPIResponse)
	},
}

// GetTaobaoFenxiaoLoginUserGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoLoginUserGetAPIResponse
func GetTaobaoFenxiaoLoginUserGetAPIResponse() *TaobaoFenxiaoLoginUserGetAPIResponse {
	return poolTaobaoFenxiaoLoginUserGetAPIResponse.Get().(*TaobaoFenxiaoLoginUserGetAPIResponse)
}

// ReleaseTaobaoFenxiaoLoginUserGetAPIResponse 将 TaobaoFenxiaoLoginUserGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoLoginUserGetAPIResponse(v *TaobaoFenxiaoLoginUserGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoLoginUserGetAPIResponse.Put(v)
}
