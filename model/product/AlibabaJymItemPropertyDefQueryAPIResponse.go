package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitempropertydefqueryAPIResponse 交易猫商品属性定义查询 API返回值
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
type AlibabajymitempropertydefqueryAPIResponse struct {
	model.CommonResponse
	AlibabajymitempropertydefqueryAPIResponseModel
}

// AlibabajymitempropertydefqueryAPIResponseModel is 交易猫商品属性定义查询 成功返回结果
type AlibabajymitempropertydefqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_property_def_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品发布填写资料定义DTO
	Result *GoodsPublishPropertyDefDetailDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
