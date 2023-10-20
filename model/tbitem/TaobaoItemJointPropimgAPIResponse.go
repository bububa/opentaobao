package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemJointPropimgAPIResponse 商品关联属性图 API返回值
// taobao.item.joint.propimg
//
// * 关联一张商品属性图片到num_iid指定的商品中&lt;br/&gt;    * 传入的num_iid所对应的商品必须属于当前会话的用户&lt;br/&gt;    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的&lt;br/&gt;    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行&lt;br/&gt;    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoItemJointPropimgAPIResponse struct {
	model.CommonResponse
	TaobaoItemJointPropimgAPIResponseModel
}

// TaobaoItemJointPropimgAPIResponseModel is 商品关联属性图 成功返回结果
type TaobaoItemJointPropimgAPIResponseModel struct {
	XMLName xml.Name `xml:"item_joint_propimg_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 属性图片对象信息
	PropImg *PropImg `json:"prop_img,omitempty" xml:"prop_img,omitempty"`
}
