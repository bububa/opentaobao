package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallDeviceStoreFollowurlGetAPIResponse 获取店铺关注链接 API返回值
// tmall.device.store.followurl.get
//
// 获取智能硬件上的关注店铺的URL
type TmallDeviceStoreFollowurlGetAPIResponse struct {
	model.CommonResponse
	TmallDeviceStoreFollowurlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallDeviceStoreFollowurlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallDeviceStoreFollowurlGetAPIResponseModel).Reset()
}

// TmallDeviceStoreFollowurlGetAPIResponseModel is 获取店铺关注链接 成功返回结果
type TmallDeviceStoreFollowurlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_device_store_followurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码短链接地址
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 二维码图片URL
	ShortImgUrl string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`
}

// Reset 清空结构体
func (m *TmallDeviceStoreFollowurlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ShortUrl = ""
	m.ShortImgUrl = ""
}

var poolTmallDeviceStoreFollowurlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallDeviceStoreFollowurlGetAPIResponse)
	},
}

// GetTmallDeviceStoreFollowurlGetAPIResponse 从 sync.Pool 获取 TmallDeviceStoreFollowurlGetAPIResponse
func GetTmallDeviceStoreFollowurlGetAPIResponse() *TmallDeviceStoreFollowurlGetAPIResponse {
	return poolTmallDeviceStoreFollowurlGetAPIResponse.Get().(*TmallDeviceStoreFollowurlGetAPIResponse)
}

// ReleaseTmallDeviceStoreFollowurlGetAPIResponse 将 TmallDeviceStoreFollowurlGetAPIResponse 保存到 sync.Pool
func ReleaseTmallDeviceStoreFollowurlGetAPIResponse(v *TmallDeviceStoreFollowurlGetAPIResponse) {
	v.Reset()
	poolTmallDeviceStoreFollowurlGetAPIResponse.Put(v)
}
