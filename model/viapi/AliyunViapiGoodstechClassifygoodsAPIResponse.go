package viapi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapigoodstechclassifygoodsAPIResponse 商品分类 API返回值
// aliyun.viapi.goodstech.classifygoods
//
// 可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapigoodstechclassifygoodsAPIResponse struct {
	model.CommonResponse
	AliyunviapigoodstechclassifygoodsAPIResponseModel
}

// AliyunviapigoodstechclassifygoodsAPIResponseModel is 商品分类 成功返回结果
type AliyunviapigoodstechclassifygoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_goodstech_classifygoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunviapigoodstechclassifygoodsData `json:"data,omitempty" xml:"data,omitempty"`
}
