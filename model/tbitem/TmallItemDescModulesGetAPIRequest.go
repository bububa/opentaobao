package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemdescmodulesgetAPIRequest 商品描述模块信息获取 API请求
// tmall.item.desc.modules.get
//
// 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallitemdescmodulesgetAPIRequest struct {
	model.Params
	// 商家主帐号id
	_usrId string
	// 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122
	_catId int64
}

// NewTmallitemdescmodulesgetRequest 初始化TmallitemdescmodulesgetAPIRequest对象
func NewTmallitemdescmodulesgetRequest() *TmallitemdescmodulesgetAPIRequest {
	return &TmallitemdescmodulesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemdescmodulesgetAPIRequest) GetApiMethodName() string {
	return "tmall.item.desc.modules.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemdescmodulesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemdescmodulesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUsrId is UsrId Setter
// 商家主帐号id
func (r *TmallitemdescmodulesgetAPIRequest) SetUsrId(_usrId string) error {
	r._usrId = _usrId
	r.Set("usr_id", _usrId)
	return nil
}

// GetUsrId UsrId Getter
func (r TmallitemdescmodulesgetAPIRequest) GetUsrId() string {
	return r._usrId
}

// SetCatId is CatId Setter
// 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&amp;path=cid:3-apiId:122
func (r *TmallitemdescmodulesgetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TmallitemdescmodulesgetAPIRequest) GetCatId() int64 {
	return r._catId
}
