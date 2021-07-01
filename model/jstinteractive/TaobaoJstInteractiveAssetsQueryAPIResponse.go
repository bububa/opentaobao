package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可配置任务素材接口 API返回值 
taobao.jst.interactive.assets.query

查询可配置任务素材列表，用以配置任务素材
*/
type TaobaoJstInteractiveAssetsQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveAssetsQueryAPIResponseModel
}

// 查询可配置任务素材接口 成功返回结果
type TaobaoJstInteractiveAssetsQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_interactive_assets_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 可配置任务素材列表
    AssetsList   []AssetsConfig `json:"assets_list,omitempty" xml:"assets_list>assets_config,omitempty"`
}
