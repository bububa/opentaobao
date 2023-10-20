package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemgameseverqueryAPIResponse 查询商品发布客户端下可用服务器列表 API返回值
// alibaba.jym.item.game.sever.query
//
// 查询商品发布客户端下可用服务器列表
type AlibabajymitemgameseverqueryAPIResponse struct {
	model.CommonResponse
	AlibabajymitemgameseverqueryAPIResponseModel
}

// AlibabajymitemgameseverqueryAPIResponseModel is 查询商品发布客户端下可用服务器列表 成功返回结果
type AlibabajymitemgameseverqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_game_sever_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 游戏服务器对象
	Result []GoodsServerDto `json:"result,omitempty" xml:"result>goods_server_dto,omitempty"`
	// 响应码信息描述
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 响应码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 是否请求成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
