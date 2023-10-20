package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalldevicestorefollowurlgetAPIResponse 获取店铺关注链接 API返回值
// tmall.device.store.followurl.get
//
// 获取智能硬件上的关注店铺的URL
type TmalldevicestorefollowurlgetAPIResponse struct {
	model.CommonResponse
	TmalldevicestorefollowurlgetAPIResponseModel
}

// TmalldevicestorefollowurlgetAPIResponseModel is 获取店铺关注链接 成功返回结果
type TmalldevicestorefollowurlgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_device_store_followurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码短链接地址
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 二维码图片URL
	ShortImgUrl string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`
}
