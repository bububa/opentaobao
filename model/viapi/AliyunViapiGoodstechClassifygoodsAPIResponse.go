package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechClassifygoodsAPIResponse 商品分类 API返回值
// aliyun.viapi.goodstech.classifygoods
//
// 可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechClassifygoodsAPIResponse struct {
	model.CommonResponse
	AliyunViapiGoodstechClassifygoodsAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechClassifygoodsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiGoodstechClassifygoodsAPIResponseModel).Reset()
}

// AliyunViapiGoodstechClassifygoodsAPIResponseModel is 商品分类 成功返回结果
type AliyunViapiGoodstechClassifygoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_goodstech_classifygoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiGoodstechClassifygoodsData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechClassifygoodsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiGoodstechClassifygoodsAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechClassifygoodsAPIResponse)
	},
}

// GetAliyunViapiGoodstechClassifygoodsAPIResponse 从 sync.Pool 获取 AliyunViapiGoodstechClassifygoodsAPIResponse
func GetAliyunViapiGoodstechClassifygoodsAPIResponse() *AliyunViapiGoodstechClassifygoodsAPIResponse {
	return poolAliyunViapiGoodstechClassifygoodsAPIResponse.Get().(*AliyunViapiGoodstechClassifygoodsAPIResponse)
}

// ReleaseAliyunViapiGoodstechClassifygoodsAPIResponse 将 AliyunViapiGoodstechClassifygoodsAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiGoodstechClassifygoodsAPIResponse(v *AliyunViapiGoodstechClassifygoodsAPIResponse) {
	v.Reset()
	poolAliyunViapiGoodstechClassifygoodsAPIResponse.Put(v)
}
