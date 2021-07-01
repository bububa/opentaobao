package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemJointImgAPIRequest
商品关联子图 API请求
taobao.item.joint.img

* 关联一张商品图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额 */
type TaobaoItemJointImgAPIRequest struct {
	model.Params
	// 商品图片id(如果是更新图片，则需要传该参数)
	_id int64
	// 商品数字ID，必选
	_numIid int64
	// 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
	_picPath string
	// 上传的图片是否关联为商品主图
	_isMajor bool
	// 图片序号
	_position int64
	// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
	_isRectangle bool
}

// New
