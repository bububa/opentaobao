package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayWordpackageGetAPIResponse 获取词包列表 API返回值
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
type TaobaoSubwayWordpackageGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayWordpackageGetAPIResponseModel
}

// TaobaoSubwayWordpackageGetAPIResponseModel is 获取词包列表 成功返回结果
type TaobaoSubwayWordpackageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_wordpackage_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词包结果列表
	ResultList []SiriusItemWordPackageDto `json:"result_list,omitempty" xml:"result_list>sirius_item_word_package_dto,omitempty"`
}
