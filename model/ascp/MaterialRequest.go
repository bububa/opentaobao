package ascp

import (
	"sync"
)

// MaterialRequest 结构体
type MaterialRequest struct {
	// 【创建时必传】图片URL [建议不超过5个图片，第一个为主图，每个URL不超过256字节]
	ImgUrl []string `json:"img_url,omitempty" xml:"img_url>string,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 【必传】店铺商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 【创建时必传】文描商品标题 [不超过32个中文字符]
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 【创建时必传】商品描述富文本
	ItemDescribeText string `json:"item_describe_text,omitempty" xml:"item_describe_text,omitempty"`
	// 【创建时必传】移动端商品描述富文本
	MobileItemDescribeText string `json:"mobile_item_describe_text,omitempty" xml:"mobile_item_describe_text,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolMaterialRequest = sync.Pool{
	New: func() any {
		return new(MaterialRequest)
	},
}

// GetMaterialRequest() 从对象池中获取MaterialRequest
func GetMaterialRequest() *MaterialRequest {
	return poolMaterialRequest.Get().(*MaterialRequest)
}

// ReleaseMaterialRequest 释放MaterialRequest
func ReleaseMaterialRequest(v *MaterialRequest) {
	v.ImgUrl = v.ImgUrl[:0]
	v.RequestId = ""
	v.ItemId = ""
	v.Title = ""
	v.ItemDescribeText = ""
	v.MobileItemDescribeText = ""
	v.RequestTime = 0
	poolMaterialRequest.Put(v)
}
