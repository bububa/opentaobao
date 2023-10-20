package wirelessshare

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessShareTpwdQueryAPIResponse 查询解析淘口令 API返回值
// taobao.wireless.share.tpwd.query
//
// 查询解析淘口令
type TaobaoWirelessShareTpwdQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWirelessShareTpwdQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWirelessShareTpwdQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWirelessShareTpwdQueryAPIResponseModel).Reset()
}

// TaobaoWirelessShareTpwdQueryAPIResponseModel is 查询解析淘口令 成功返回结果
type TaobaoWirelessShareTpwdQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wireless_share_tpwd_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘口令-文案
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 淘口令-宝贝
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 如果是宝贝，则为宝贝价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 跳转url(长链)
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// nativeUrl
	NativeUrl string `json:"native_url,omitempty" xml:"native_url,omitempty"`
	// thumbPicUrl
	ThumbPicUrl string `json:"thumb_pic_url,omitempty" xml:"thumb_pic_url,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWirelessShareTpwdQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = ""
	m.Title = ""
	m.Price = ""
	m.PicUrl = ""
	m.Url = ""
	m.NativeUrl = ""
	m.ThumbPicUrl = ""
	m.Suc = false
}

var poolTaobaoWirelessShareTpwdQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWirelessShareTpwdQueryAPIResponse)
	},
}

// GetTaobaoWirelessShareTpwdQueryAPIResponse 从 sync.Pool 获取 TaobaoWirelessShareTpwdQueryAPIResponse
func GetTaobaoWirelessShareTpwdQueryAPIResponse() *TaobaoWirelessShareTpwdQueryAPIResponse {
	return poolTaobaoWirelessShareTpwdQueryAPIResponse.Get().(*TaobaoWirelessShareTpwdQueryAPIResponse)
}

// ReleaseTaobaoWirelessShareTpwdQueryAPIResponse 将 TaobaoWirelessShareTpwdQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWirelessShareTpwdQueryAPIResponse(v *TaobaoWirelessShareTpwdQueryAPIResponse) {
	v.Reset()
	poolTaobaoWirelessShareTpwdQueryAPIResponse.Put(v)
}
