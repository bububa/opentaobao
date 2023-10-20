package jstinteractive

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveassetsconfiguredqueryAPIResponse 查询已配置的任务素材列表接口 API返回值
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
type TaobaojstinteractiveassetsconfiguredqueryAPIResponse struct {
	model.CommonResponse
	TaobaojstinteractiveassetsconfiguredqueryAPIResponseModel
}

// TaobaojstinteractiveassetsconfiguredqueryAPIResponseModel is 查询已配置的任务素材列表接口 成功返回结果
type TaobaojstinteractiveassetsconfiguredqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_assets_configured_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 可配置任务素材列表
	AssetsList []AssetsConfig `json:"assets_list,omitempty" xml:"assets_list>assets_config,omitempty"`
}
