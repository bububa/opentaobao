package tmallcms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse 创建流量宝活动链接 API返回值
// tmall.marketing.liuliangbao.spreadlink.create
//
// 通过源活动链接和商家ID，创建流量宝活动链接
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse struct {
	model.CommonResponse
	TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponseModel).Reset()
}

// TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponseModel is 创建流量宝活动链接 成功返回结果
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_marketing_liuliangbao_spreadlink_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 流量宝系统执行结果
	Llbresult *LLBApiResult `json:"llbresult,omitempty" xml:"llbresult,omitempty"`
}

// Reset 清空结构体
func (m *TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Llbresult = nil
}

var poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse)
	},
}

// GetTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse 从 sync.Pool 获取 TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse
func GetTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse() *TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse {
	return poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse.Get().(*TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse)
}

// ReleaseTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse 将 TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse(v *TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse) {
	v.Reset()
	poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse.Put(v)
}
