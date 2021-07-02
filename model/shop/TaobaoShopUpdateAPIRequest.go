package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoShopUpdateAPIRequest 更新店铺基本信息 API请求
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
type TaobaoShopUpdateAPIRequest struct {
	model.Params
	// 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
	_title string
	// 店铺公告。不超过1024个字符
	_bulletin string
	// 店铺描述。10～2000个字符以内
	_desc string
}

// NewTaobaoShopUpdateRequest 初始化TaobaoShopUpdateAPIRequest对象
func NewTaobaoShopUpdateRequest() *TaobaoShopUpdateAPIRequest {
	return &TaobaoShopUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoShopUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.shop.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoShopUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTitle is Title Setter
// 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
func (r *TaobaoShopUpdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoShopUpdateAPIRequest) GetTitle() string {
	return r._title
}

// SetBulletin is Bulletin Setter
// 店铺公告。不超过1024个字符
func (r *TaobaoShopUpdateAPIRequest) SetBulletin(_bulletin string) error {
	r._bulletin = _bulletin
	r.Set("bulletin", _bulletin)
	return nil
}

// GetBulletin Bulletin Getter
func (r TaobaoShopUpdateAPIRequest) GetBulletin() string {
	return r._bulletin
}

// SetDesc is Desc Setter
// 店铺描述。10～2000个字符以内
func (r *TaobaoShopUpdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoShopUpdateAPIRequest) GetDesc() string {
	return r._desc
}
