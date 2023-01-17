package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsDummySendAPIRequest 无需物流（虚拟）发货处理 API请求
// taobao.logistics.dummy.send
//
// 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
type TaobaoLogisticsDummySendAPIRequest struct {
	model.Params
	// feature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
	_feature string
	// 商家的IP地址
	_sellerIp string
	// 淘宝交易ID
	_tid int64
}

// NewTaobaoLogisticsDummySendRequest 初始化TaobaoLogisticsDummySendAPIRequest对象
func NewTaobaoLogisticsDummySendRequest() *TaobaoLogisticsDummySendAPIRequest {
	return &TaobaoLogisticsDummySendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsDummySendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.dummy.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsDummySendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsDummySendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeature is Feature Setter
// feature参数格式&lt;br&gt;范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B&lt;br&gt;identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔&lt;br&gt;“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。&#34;|&#34;不同商品间的分隔符。&lt;br&gt;例1商品和2商品，之间就用&#34;|&#34;分开。&lt;br&gt;TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)&lt;br&gt;&#34;:&#34;TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。&lt;br&gt;&#34;,&#34; 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。&lt;br&gt;具体:当订单中A商品的数量为2个，其中手机串号分别为&#34;12345&#34;,&#34;67890&#34;。&lt;br&gt;参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的&#34;12345&#34;,&#34;67890&#34;.说明本订单A宝贝的数量为2，值分别为&#34;12345&#34;,&#34;67890&#34;。&lt;br&gt;当存在&#34;|&#34;时，就说明订单中存在多个商品，商品间用&#34;|&#34;分隔了开来。|&#34;之后的内容含义同上。
func (r *TaobaoLogisticsDummySendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r TaobaoLogisticsDummySendAPIRequest) GetFeature() string {
	return r._feature
}

// SetSellerIp is SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsDummySendAPIRequest) SetSellerIp(_sellerIp string) error {
	r._sellerIp = _sellerIp
	r.Set("seller_ip", _sellerIp)
	return nil
}

// GetSellerIp SellerIp Getter
func (r TaobaoLogisticsDummySendAPIRequest) GetSellerIp() string {
	return r._sellerIp
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsDummySendAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoLogisticsDummySendAPIRequest) GetTid() int64 {
	return r._tid
}
