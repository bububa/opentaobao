package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallGoodsSyncAPIResponse 第三方商家接入采购商城-商品同步 API返回值
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
type AlibabaPurCmallGoodsSyncAPIResponse struct {
	model.CommonResponse
	AlibabaPurCmallGoodsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurCmallGoodsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurCmallGoodsSyncAPIResponseModel).Reset()
}

// AlibabaPurCmallGoodsSyncAPIResponseModel is 第三方商家接入采购商城-商品同步 成功返回结果
type AlibabaPurCmallGoodsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_cmall_goods_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurCmallGoodsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurCmallGoodsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurCmallGoodsSyncAPIResponse)
	},
}

// GetAlibabaPurCmallGoodsSyncAPIResponse 从 sync.Pool 获取 AlibabaPurCmallGoodsSyncAPIResponse
func GetAlibabaPurCmallGoodsSyncAPIResponse() *AlibabaPurCmallGoodsSyncAPIResponse {
	return poolAlibabaPurCmallGoodsSyncAPIResponse.Get().(*AlibabaPurCmallGoodsSyncAPIResponse)
}

// ReleaseAlibabaPurCmallGoodsSyncAPIResponse 将 AlibabaPurCmallGoodsSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurCmallGoodsSyncAPIResponse(v *AlibabaPurCmallGoodsSyncAPIResponse) {
	v.Reset()
	poolAlibabaPurCmallGoodsSyncAPIResponse.Put(v)
}
