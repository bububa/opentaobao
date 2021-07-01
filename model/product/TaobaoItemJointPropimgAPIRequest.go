package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemJointPropimgAPIRequest
商品关联属性图 API请求
taobao.item.joint.propimg

* 关联一张商品属性图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张） */
type TaobaoItemJointPropimgAPIRequest struct {
	model.Params
	// 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
	_properties string
	// 属性图片ID。如果是新增不需要填写
	_id int64
	// 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
	_picPath string
	// 商品数字ID，必选
	_numIid int64
	// 图片序号
	_position int64
}

// New
