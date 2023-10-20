package uscesl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizBrandInsertAPIResponse 新增电子价签商家 API返回值
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
type TaobaoUsceslBizBrandInsertAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizBrandInsertAPIResponseModel
}

// TaobaoUsceslBizBrandInsertAPIResponseModel is 新增电子价签商家 成功返回结果
type TaobaoUsceslBizBrandInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_brand_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
