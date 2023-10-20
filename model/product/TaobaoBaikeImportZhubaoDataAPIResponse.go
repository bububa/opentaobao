package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaikeimportzhubaodataAPIResponse 导入数据到商品百科服务 API返回值
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
type TaobaobaikeimportzhubaodataAPIResponse struct {
	model.CommonResponse
	TaobaobaikeimportzhubaodataAPIResponseModel
}

// TaobaobaikeimportzhubaodataAPIResponseModel is 导入数据到商品百科服务 成功返回结果
type TaobaobaikeimportzhubaodataAPIResponseModel struct {
	XMLName xml.Name `xml:"baike_import_zhubao_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaobaikeimportzhubaodataResult `json:"result,omitempty" xml:"result,omitempty"`
}
