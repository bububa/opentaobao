package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemstoreBandingAPIResponse 商品关联绑定接口 API返回值
// taobao.qimen.itemstore.banding
//
// 商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
type TaobaoQimenItemstoreBandingAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemstoreBandingAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemstoreBandingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemstoreBandingAPIResponseModel).Reset()
}

// TaobaoQimenItemstoreBandingAPIResponseModel is 商品关联绑定接口 成功返回结果
type TaobaoQimenItemstoreBandingAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemstore_banding_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应描述
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemstoreBandingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Flag = ""
	m.QimenCode = ""
}

var poolTaobaoQimenItemstoreBandingAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemstoreBandingAPIResponse)
	},
}

// GetTaobaoQimenItemstoreBandingAPIResponse 从 sync.Pool 获取 TaobaoQimenItemstoreBandingAPIResponse
func GetTaobaoQimenItemstoreBandingAPIResponse() *TaobaoQimenItemstoreBandingAPIResponse {
	return poolTaobaoQimenItemstoreBandingAPIResponse.Get().(*TaobaoQimenItemstoreBandingAPIResponse)
}

// ReleaseTaobaoQimenItemstoreBandingAPIResponse 将 TaobaoQimenItemstoreBandingAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemstoreBandingAPIResponse(v *TaobaoQimenItemstoreBandingAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemstoreBandingAPIResponse.Put(v)
}
