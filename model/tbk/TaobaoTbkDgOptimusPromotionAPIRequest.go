package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgoptimuspromotionAPIRequest 淘宝客-推广者-权益物料精选 API请求
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaotbkdgoptimuspromotionAPIRequest struct {
	model.Params
	// 页大小，一次请求请限制在10以内
	_pagesize int64
	// 第几页，默认：1
	_pagenum int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneid int64
	// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
	_promotionid int64
}

// NewTaobaotbkdgoptimuspromotionRequest 初始化TaobaotbkdgoptimuspromotionAPIRequest对象
func NewTaobaotbkdgoptimuspromotionRequest() *TaobaotbkdgoptimuspromotionAPIRequest {
	return &TaobaotbkdgoptimuspromotionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.optimus.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPagesize is Pagesize Setter
// 页大小，一次请求请限制在10以内
func (r *TaobaotbkdgoptimuspromotionAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPagenum is Pagenum Setter
// 第几页，默认：1
func (r *TaobaotbkdgoptimuspromotionAPIRequest) SetPagenum(_pagenum int64) error {
	r._pagenum = _pagenum
	r.Set("page_num", _pagenum)
	return nil
}

// GetPagenum Pagenum Getter
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetPagenum() int64 {
	return r._pagenum
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkdgoptimuspromotionAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetPromotionid is Promotionid Setter
// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
func (r *TaobaotbkdgoptimuspromotionAPIRequest) SetPromotionid(_promotionid int64) error {
	r._promotionid = _promotionid
	r.Set("promotion_id", _promotionid)
	return nil
}

// GetPromotionid Promotionid Getter
func (r TaobaotbkdgoptimuspromotionAPIRequest) GetPromotionid() int64 {
	return r._promotionid
}
