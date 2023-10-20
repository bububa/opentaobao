package caipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoGoodsInfoGetAPIResponse 根据卖家id与appkey获取商品信息 API返回值
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
type TaobaoCaipiaoGoodsInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoGoodsInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCaipiaoGoodsInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCaipiaoGoodsInfoGetAPIResponseModel).Reset()
}

// TaobaoCaipiaoGoodsInfoGetAPIResponseModel is 根据卖家id与appkey获取商品信息 成功返回结果
type TaobaoCaipiaoGoodsInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_goods_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询的结果列表
	Results []LotteryWangcaiSellerGoodsInfo `json:"results,omitempty" xml:"results>lottery_wangcai_seller_goods_info,omitempty"`
	// 返回列表的大小
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCaipiaoGoodsInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalResults = 0
}

var poolTaobaoCaipiaoGoodsInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCaipiaoGoodsInfoGetAPIResponse)
	},
}

// GetTaobaoCaipiaoGoodsInfoGetAPIResponse 从 sync.Pool 获取 TaobaoCaipiaoGoodsInfoGetAPIResponse
func GetTaobaoCaipiaoGoodsInfoGetAPIResponse() *TaobaoCaipiaoGoodsInfoGetAPIResponse {
	return poolTaobaoCaipiaoGoodsInfoGetAPIResponse.Get().(*TaobaoCaipiaoGoodsInfoGetAPIResponse)
}

// ReleaseTaobaoCaipiaoGoodsInfoGetAPIResponse 将 TaobaoCaipiaoGoodsInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCaipiaoGoodsInfoGetAPIResponse(v *TaobaoCaipiaoGoodsInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoCaipiaoGoodsInfoGetAPIResponse.Put(v)
}
