package uscesl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizBrandInsertAPIRequest 新增电子价签商家 API请求
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
type TaobaoUsceslBizBrandInsertAPIRequest struct {
	model.Params
	// 商家名称
	_brandName string
	// 商家外部编号
	_brandOutCode string
}

// NewTaobaoUsceslBizBrandInsertRequest 初始化TaobaoUsceslBizBrandInsertAPIRequest对象
func NewTaobaoUsceslBizBrandInsertRequest() *TaobaoUsceslBizBrandInsertAPIRequest {
	return &TaobaoUsceslBizBrandInsertAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizBrandInsertAPIRequest) Reset() {
	r._brandName = ""
	r._brandOutCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizBrandInsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.brand.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizBrandInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizBrandInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandName is BrandName Setter
// 商家名称
func (r *TaobaoUsceslBizBrandInsertAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaoUsceslBizBrandInsertAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetBrandOutCode is BrandOutCode Setter
// 商家外部编号
func (r *TaobaoUsceslBizBrandInsertAPIRequest) SetBrandOutCode(_brandOutCode string) error {
	r._brandOutCode = _brandOutCode
	r.Set("brand_out_code", _brandOutCode)
	return nil
}

// GetBrandOutCode BrandOutCode Getter
func (r TaobaoUsceslBizBrandInsertAPIRequest) GetBrandOutCode() string {
	return r._brandOutCode
}

var poolTaobaoUsceslBizBrandInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizBrandInsertRequest()
	},
}

// GetTaobaoUsceslBizBrandInsertRequest 从 sync.Pool 获取 TaobaoUsceslBizBrandInsertAPIRequest
func GetTaobaoUsceslBizBrandInsertAPIRequest() *TaobaoUsceslBizBrandInsertAPIRequest {
	return poolTaobaoUsceslBizBrandInsertAPIRequest.Get().(*TaobaoUsceslBizBrandInsertAPIRequest)
}

// ReleaseTaobaoUsceslBizBrandInsertAPIRequest 将 TaobaoUsceslBizBrandInsertAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizBrandInsertAPIRequest(v *TaobaoUsceslBizBrandInsertAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizBrandInsertAPIRequest.Put(v)
}
