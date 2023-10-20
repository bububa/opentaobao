package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponse 基于游戏id临时上下架智能发布入口 API返回值
// alibaba.jym.goods.external.goods.vmos.offon.game
//
// 基于游戏id临时上下架智能发布入口
type AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponse struct {
	model.CommonResponse
	AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponseModel
}

// AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponseModel is 基于游戏id临时上下架智能发布入口 成功返回结果
type AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_goods_external_goods_vmos_offon_game_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 调用结果描述
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 代表调用成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
