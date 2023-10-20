package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemJointImgAPIResponse 商品关联子图 API返回值
// taobao.item.joint.img
//
// * 关联一张商品图片到num_iid指定的商品中&lt;br/&gt;    * 传入的num_iid所对应的商品必须属于当前会话的用户&lt;br/&gt;    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行&lt;br/&gt;    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgAPIResponse struct {
	model.CommonResponse
	TaobaoItemJointImgAPIResponseModel
}

// TaobaoItemJointImgAPIResponseModel is 商品关联子图 成功返回结果
type TaobaoItemJointImgAPIResponseModel struct {
	XMLName xml.Name `xml:"item_joint_img_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品图片信息
	ItemImg *ItemImg `json:"item_img,omitempty" xml:"item_img,omitempty"`
}
