package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateQueryumpactidAPIRequest 通过券模板查询券活动id接口 API请求
// alibaba.wdk.coupon.template.queryumpactid
//
// 当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
type AlibabaWdkCouponTemplateQueryumpactidAPIRequest struct {
	model.Params
	// 券模版id列表
	_sourceIds []string
	// 优惠券类型
	_wdkCouponType int64
}

// NewAlibabaWdkCouponTemplateQueryumpactidRequest 初始化AlibabaWdkCouponTemplateQueryumpactidAPIRequest对象
func NewAlibabaWdkCouponTemplateQueryumpactidRequest() *AlibabaWdkCouponTemplateQueryumpactidAPIRequest {
	return &AlibabaWdkCouponTemplateQueryumpactidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) Reset() {
	r._sourceIds = r._sourceIds[:0]
	r._wdkCouponType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.coupon.template.queryumpactid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceIds is SourceIds Setter
// 券模版id列表
func (r *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) SetSourceIds(_sourceIds []string) error {
	r._sourceIds = _sourceIds
	r.Set("source_ids", _sourceIds)
	return nil
}

// GetSourceIds SourceIds Getter
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetSourceIds() []string {
	return r._sourceIds
}

// SetWdkCouponType is WdkCouponType Setter
// 优惠券类型
func (r *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) SetWdkCouponType(_wdkCouponType int64) error {
	r._wdkCouponType = _wdkCouponType
	r.Set("wdk_coupon_type", _wdkCouponType)
	return nil
}

// GetWdkCouponType WdkCouponType Getter
func (r AlibabaWdkCouponTemplateQueryumpactidAPIRequest) GetWdkCouponType() int64 {
	return r._wdkCouponType
}

var poolAlibabaWdkCouponTemplateQueryumpactidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkCouponTemplateQueryumpactidRequest()
	},
}

// GetAlibabaWdkCouponTemplateQueryumpactidRequest 从 sync.Pool 获取 AlibabaWdkCouponTemplateQueryumpactidAPIRequest
func GetAlibabaWdkCouponTemplateQueryumpactidAPIRequest() *AlibabaWdkCouponTemplateQueryumpactidAPIRequest {
	return poolAlibabaWdkCouponTemplateQueryumpactidAPIRequest.Get().(*AlibabaWdkCouponTemplateQueryumpactidAPIRequest)
}

// ReleaseAlibabaWdkCouponTemplateQueryumpactidAPIRequest 将 AlibabaWdkCouponTemplateQueryumpactidAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkCouponTemplateQueryumpactidAPIRequest(v *AlibabaWdkCouponTemplateQueryumpactidAPIRequest) {
	v.Reset()
	poolAlibabaWdkCouponTemplateQueryumpactidAPIRequest.Put(v)
}
