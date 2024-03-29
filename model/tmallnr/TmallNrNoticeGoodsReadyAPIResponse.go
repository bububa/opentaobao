package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrNoticeGoodsReadyAPIResponse 同步天猫配送人员信息 API返回值
// tmall.nr.notice.goods.ready
//
// 接收商家的配送人员信息，和第三公司信息及提货码
type TmallNrNoticeGoodsReadyAPIResponse struct {
	model.CommonResponse
	TmallNrNoticeGoodsReadyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrNoticeGoodsReadyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrNoticeGoodsReadyAPIResponseModel).Reset()
}

// TmallNrNoticeGoodsReadyAPIResponseModel is 同步天猫配送人员信息 成功返回结果
type TmallNrNoticeGoodsReadyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_notice_goods_ready_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode2 string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`
	// resultData
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrNoticeGoodsReadyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessage = ""
	m.ErrorCode2 = ""
	m.ResultData = false
	m.IsSuccess = false
}

var poolTmallNrNoticeGoodsReadyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrNoticeGoodsReadyAPIResponse)
	},
}

// GetTmallNrNoticeGoodsReadyAPIResponse 从 sync.Pool 获取 TmallNrNoticeGoodsReadyAPIResponse
func GetTmallNrNoticeGoodsReadyAPIResponse() *TmallNrNoticeGoodsReadyAPIResponse {
	return poolTmallNrNoticeGoodsReadyAPIResponse.Get().(*TmallNrNoticeGoodsReadyAPIResponse)
}

// ReleaseTmallNrNoticeGoodsReadyAPIResponse 将 TmallNrNoticeGoodsReadyAPIResponse 保存到 sync.Pool
func ReleaseTmallNrNoticeGoodsReadyAPIResponse(v *TmallNrNoticeGoodsReadyAPIResponse) {
	v.Reset()
	poolTmallNrNoticeGoodsReadyAPIResponse.Put(v)
}
