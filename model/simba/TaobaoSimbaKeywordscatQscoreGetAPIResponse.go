package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordscatQscoreGetAPIResponse 取得一个推广组的所有关键词和类目出价的质量得分 API返回值
// taobao.simba.keywordscat.qscore.get
//
// 取得一个推广组的所有关键词和类目出价的质量得分列表
type TaobaoSimbaKeywordscatQscoreGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordscatQscoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordscatQscoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordscatQscoreGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordscatQscoreGetAPIResponseModel is 取得一个推广组的所有关键词和类目出价的质量得分 成功返回结果
type TaobaoSimbaKeywordscatQscoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordscat_qscore_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目出价和词的质量得分对象
	Qscore *Qscore `json:"qscore,omitempty" xml:"qscore,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordscatQscoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Qscore = nil
}

var poolTaobaoSimbaKeywordscatQscoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordscatQscoreGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordscatQscoreGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordscatQscoreGetAPIResponse
func GetTaobaoSimbaKeywordscatQscoreGetAPIResponse() *TaobaoSimbaKeywordscatQscoreGetAPIResponse {
	return poolTaobaoSimbaKeywordscatQscoreGetAPIResponse.Get().(*TaobaoSimbaKeywordscatQscoreGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordscatQscoreGetAPIResponse 将 TaobaoSimbaKeywordscatQscoreGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordscatQscoreGetAPIResponse(v *TaobaoSimbaKeywordscatQscoreGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordscatQscoreGetAPIResponse.Put(v)
}
