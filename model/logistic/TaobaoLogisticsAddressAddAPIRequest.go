package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressAddAPIRequest
卖家地址库新增接口 API请求
taobao.logistics.address.add

通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库 */
type TaobaoLogisticsAddressAddAPIRequest struct {
	model.Params
	// 联系人姓名 <font color='red'>长度不可超过20个字节</font>
	_contactName string
	// 所在省
	_province string
	// 所在市
	_city string
	// 区、县<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
	_country string
	// 详细街道地址，不需要重复填写省/市/区
	_addr string
	// 地区邮政编码<br/><br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font>
	_zipCode string
	// 电话号码,手机与电话必需有一个
	_phone string
	// 手机号码，手机与电话必需有一个<br/><br><font color='red'>手机号码不能超过20位</font>
	_mobilePhone string
	// 公司名称,<br><font color="red">公司名称长能不能超过20字节</font>
	_sellerCompany string
	// 备注,<br><font color='red'>备注不能超过256字节</font>
	_memo string
	// 默认取货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font>
	_getDef bool
	// 默认退货地址。<br><br/><font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font>
	_cancelDef bool
}

// New
