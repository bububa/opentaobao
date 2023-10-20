package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarkeywordsqscoresplitgetAPIResponse (新)销量明星质量分相关接口 API返回值
// taobao.simba.salestar.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaosimbasalestarkeywordsqscoresplitgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbasalestarkeywordsqscoresplitgetAPIResponseModel
}

// TaobaosimbasalestarkeywordsqscoresplitgetAPIResponseModel is (新)销量明星质量分相关接口 成功返回结果
type TaobaosimbasalestarkeywordsqscoresplitgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_keywords_qscore_split_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosimbasalestarkeywordsqscoresplitgetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
