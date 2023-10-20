package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtNewcouponSendAPIResponse 券发放接口 API返回值
// tmall.nrt.newcoupon.send
//
// 券发放接口
type TmallNrtNewcouponSendAPIResponse struct {
	model.CommonResponse
	TmallNrtNewcouponSendAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtNewcouponSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtNewcouponSendAPIResponseModel).Reset()
}

// TmallNrtNewcouponSendAPIResponseModel is 券发放接口 成功返回结果
type TmallNrtNewcouponSendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_newcoupon_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TmallNrtNewcouponSendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtNewcouponSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtNewcouponSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtNewcouponSendAPIResponse)
	},
}

// GetTmallNrtNewcouponSendAPIResponse 从 sync.Pool 获取 TmallNrtNewcouponSendAPIResponse
func GetTmallNrtNewcouponSendAPIResponse() *TmallNrtNewcouponSendAPIResponse {
	return poolTmallNrtNewcouponSendAPIResponse.Get().(*TmallNrtNewcouponSendAPIResponse)
}

// ReleaseTmallNrtNewcouponSendAPIResponse 将 TmallNrtNewcouponSendAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtNewcouponSendAPIResponse(v *TmallNrtNewcouponSendAPIResponse) {
	v.Reset()
	poolTmallNrtNewcouponSendAPIResponse.Put(v)
}
