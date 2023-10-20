package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse 查询达摩盘精选人群模板 API返回值
// taobao.onebp.dkx.crowd.crowd.finddmpcrowd
//
// 查询达摩盘精选人群模板；使用方法为先查询出topic和对应的templateId（需要一一对应），然后将想使用的模板topic&amp;templateId组填入Add接口中的new_dmp_template_crowd结构中提交即可。
type TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponseModel is 查询达摩盘精选人群模板 成功返回结果
type TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_crowd_crowd_finddmpcrowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse)
	},
}

// GetTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse
func GetTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse() *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse {
	return poolTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse.Get().(*TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse 将 TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse(v *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIResponse.Put(v)
}
