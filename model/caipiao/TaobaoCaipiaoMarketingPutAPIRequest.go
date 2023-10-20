package caipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoMarketingPutAPIRequest 创建或修改商家送彩票活动 API请求
// taobao.caipiao.marketing.put
//
// 卖家通过此接口新增或修改送彩票活动的配置，比如活动时间、活动的条件等。
//
// 店铺装修、宝贝详情页的宣导素材示例：
// https://gw.alicdn.com/tfs/TB1_nOiSXXXXXbgXXXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1FZX6SXXXXXXzXFXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1z4t8SXXXXXckXpXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1BhqgSXXXXXcDXXXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1TYt9SXXXXXXAXFXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1tzpNSXXXXXacXVXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1UXdxSXXXXXXsapXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1_gV.SXXXXXbZXpXXXXXXXXXX-790-280.png
type TaobaoCaipiaoMarketingPutAPIRequest struct {
	model.Params
	// 活动详情设置
	_detail *WangcaiMarketingDetail
}

// NewTaobaoCaipiaoMarketingPutRequest 初始化TaobaoCaipiaoMarketingPutAPIRequest对象
func NewTaobaoCaipiaoMarketingPutRequest() *TaobaoCaipiaoMarketingPutAPIRequest {
	return &TaobaoCaipiaoMarketingPutAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCaipiaoMarketingPutAPIRequest) Reset() {
	r._detail = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.marketing.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetail is Detail Setter
// 活动详情设置
func (r *TaobaoCaipiaoMarketingPutAPIRequest) SetDetail(_detail *WangcaiMarketingDetail) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// GetDetail Detail Getter
func (r TaobaoCaipiaoMarketingPutAPIRequest) GetDetail() *WangcaiMarketingDetail {
	return r._detail
}

var poolTaobaoCaipiaoMarketingPutAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCaipiaoMarketingPutRequest()
	},
}

// GetTaobaoCaipiaoMarketingPutRequest 从 sync.Pool 获取 TaobaoCaipiaoMarketingPutAPIRequest
func GetTaobaoCaipiaoMarketingPutAPIRequest() *TaobaoCaipiaoMarketingPutAPIRequest {
	return poolTaobaoCaipiaoMarketingPutAPIRequest.Get().(*TaobaoCaipiaoMarketingPutAPIRequest)
}

// ReleaseTaobaoCaipiaoMarketingPutAPIRequest 将 TaobaoCaipiaoMarketingPutAPIRequest 放入 sync.Pool
func ReleaseTaobaoCaipiaoMarketingPutAPIRequest(v *TaobaoCaipiaoMarketingPutAPIRequest) {
	v.Reset()
	poolTaobaoCaipiaoMarketingPutAPIRequest.Put(v)
}
