package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemDescModulesGetAPIRequest
商品描述模块信息获取 API请求
tmall.item.desc.modules.get

商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。 */
type TmallItemDescModulesGetAPIRequest struct {
	model.Params
	// 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122
	_catId int64
	// 商家主帐号id
	_usrId string
}

// New
