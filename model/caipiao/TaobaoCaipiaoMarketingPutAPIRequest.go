package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaomarketingputAPIRequest 创建或修改商家送彩票活动 API请求
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
type TaobaocaipiaomarketingputAPIRequest struct {
	model.Params
	// 活动详情设置
	_detail *WangcaiMarketingDetail
}

// NewTaobaocaipiaomarketingputRequest 初始化TaobaocaipiaomarketingputAPIRequest对象
func NewTaobaocaipiaomarketingputRequest() *TaobaocaipiaomarketingputAPIRequest {
	return &TaobaocaipiaomarketingputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaomarketingputAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.marketing.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaomarketingputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaomarketingputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetail is Detail Setter
// 活动详情设置
func (r *TaobaocaipiaomarketingputAPIRequest) SetDetail(_detail *WangcaiMarketingDetail) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// GetDetail Detail Getter
func (r TaobaocaipiaomarketingputAPIRequest) GetDetail() *WangcaiMarketingDetail {
	return r._detail
}
