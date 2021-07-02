package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtSimpleitemQueryAPIResponse 简易商品查询接口 API返回值
// tmall.nrt.simpleitem.query
//
// 为居然之家和阿里的合资公司 homeStyler提供简易的商品信息查询 包含商品名称  图片 状态
//
// 后续合资公司服务会迁移到内网 暂时过渡用
type TmallNrtSimpleitemQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtSimpleitemQueryAPIResponseModel
}

// TmallNrtSimpleitemQueryAPIResponseModel is 简易商品查询接口 成功返回结果
type TmallNrtSimpleitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_simpleitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 对象
	Models []TmallNrtSimpleitemQueryModel `json:"models,omitempty" xml:"models>tmall_nrt_simpleitem_query_model,omitempty"`
}
