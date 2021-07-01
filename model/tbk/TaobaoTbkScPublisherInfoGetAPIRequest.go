package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkScPublisherInfoGetAPIRequest
淘宝客-公用-私域用户备案信息查询 API请求
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。 */
type TaobaoTbkScPublisherInfoGetAPIRequest struct {
	model.Params
	// 类型，必选 1:渠道信息；2:会员信息
	_infoType int64
	// 渠道独占 - 渠道关系ID
	_relationId int64
	// 第几页
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
	_relationApp string
	// 会员独占 - 会员运营ID
	_specialId string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalId string
	// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
	_externalType int64
}

// New
