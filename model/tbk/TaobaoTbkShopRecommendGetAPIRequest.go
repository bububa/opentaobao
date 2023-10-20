package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkshoprecommendgetAPIRequest 淘宝客-公用-店铺关联推荐 API请求
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
type TaobaotbkshoprecommendgetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 返回数量，默认20，最大值40
	_count int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 卖家Id
	_userid int64
}

// NewTaobaotbkshoprecommendgetRequest 初始化TaobaotbkshoprecommendgetAPIRequest对象
func NewTaobaotbkshoprecommendgetRequest() *TaobaotbkshoprecommendgetAPIRequest {
	return &TaobaotbkshoprecommendgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkshoprecommendgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.shop.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkshoprecommendgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkshoprecommendgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaotbkshoprecommendgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotbkshoprecommendgetAPIRequest) GetFields() string {
	return r._fields
}

// SetCount is Count Setter
// 返回数量，默认20，最大值40
func (r *TaobaotbkshoprecommendgetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaotbkshoprecommendgetAPIRequest) GetCount() int64 {
	return r._count
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaotbkshoprecommendgetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaotbkshoprecommendgetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetUserid is Userid Setter
// 卖家Id
func (r *TaobaotbkshoprecommendgetAPIRequest) SetUserid(_userid int64) error {
	r._userid = _userid
	r.Set("user_id", _userid)
	return nil
}

// GetUserid Userid Getter
func (r TaobaotbkshoprecommendgetAPIRequest) GetUserid() int64 {
	return r._userid
}
