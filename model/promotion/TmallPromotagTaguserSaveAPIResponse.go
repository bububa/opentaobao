package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserSaveAPIResponse 给用户打上优惠标签 API返回值
// tmall.promotag.taguser.save
//
// 给用户载体打标
type TmallPromotagTaguserSaveAPIResponse struct {
	model.CommonResponse
	TmallPromotagTaguserSaveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotagTaguserSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotagTaguserSaveAPIResponseModel).Reset()
}

// TmallPromotagTaguserSaveAPIResponseModel is 给用户打上优惠标签 成功返回结果
type TmallPromotagTaguserSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_taguser_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打标结果是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotagTaguserSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTmallPromotagTaguserSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotagTaguserSaveAPIResponse)
	},
}

// GetTmallPromotagTaguserSaveAPIResponse 从 sync.Pool 获取 TmallPromotagTaguserSaveAPIResponse
func GetTmallPromotagTaguserSaveAPIResponse() *TmallPromotagTaguserSaveAPIResponse {
	return poolTmallPromotagTaguserSaveAPIResponse.Get().(*TmallPromotagTaguserSaveAPIResponse)
}

// ReleaseTmallPromotagTaguserSaveAPIResponse 将 TmallPromotagTaguserSaveAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotagTaguserSaveAPIResponse(v *TmallPromotagTaguserSaveAPIResponse) {
	v.Reset()
	poolTmallPromotagTaguserSaveAPIResponse.Put(v)
}
