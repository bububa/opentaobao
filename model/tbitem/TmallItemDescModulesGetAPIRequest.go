package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemDescModulesGetAPIRequest 商品描述模块信息获取 API请求
// tmall.item.desc.modules.get
//
// 商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
type TmallItemDescModulesGetAPIRequest struct {
	model.Params
	// 商家主帐号id
	_usrId string
	// 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122
	_catId int64
}

// NewTmallItemDescModulesGetRequest 初始化TmallItemDescModulesGetAPIRequest对象
func NewTmallItemDescModulesGetRequest() *TmallItemDescModulesGetAPIRequest {
	return &TmallItemDescModulesGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemDescModulesGetAPIRequest) Reset() {
	r._usrId = ""
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemDescModulesGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.desc.modules.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemDescModulesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemDescModulesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUsrId is UsrId Setter
// 商家主帐号id
func (r *TmallItemDescModulesGetAPIRequest) SetUsrId(_usrId string) error {
	r._usrId = _usrId
	r.Set("usr_id", _usrId)
	return nil
}

// GetUsrId UsrId Getter
func (r TmallItemDescModulesGetAPIRequest) GetUsrId() string {
	return r._usrId
}

// SetCatId is CatId Setter
// 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&amp;path=cid:3-apiId:122
func (r *TmallItemDescModulesGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TmallItemDescModulesGetAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolTmallItemDescModulesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemDescModulesGetRequest()
	},
}

// GetTmallItemDescModulesGetRequest 从 sync.Pool 获取 TmallItemDescModulesGetAPIRequest
func GetTmallItemDescModulesGetAPIRequest() *TmallItemDescModulesGetAPIRequest {
	return poolTmallItemDescModulesGetAPIRequest.Get().(*TmallItemDescModulesGetAPIRequest)
}

// ReleaseTmallItemDescModulesGetAPIRequest 将 TmallItemDescModulesGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemDescModulesGetAPIRequest(v *TmallItemDescModulesGetAPIRequest) {
	v.Reset()
	poolTmallItemDescModulesGetAPIRequest.Put(v)
}
