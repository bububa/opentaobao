package openmall

import (
	"sync"
)

// RefundMessagePic 结构体
type RefundMessagePic struct {
	// 退款单图片留言
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 退款单图片地址
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 使用taobao.openmall.refund.image.upload得到的上传token
	PicToken string `json:"pic_token,omitempty" xml:"pic_token,omitempty"`
}

var poolRefundMessagePic = sync.Pool{
	New: func() any {
		return new(RefundMessagePic)
	},
}

// GetRefundMessagePic() 从对象池中获取RefundMessagePic
func GetRefundMessagePic() *RefundMessagePic {
	return poolRefundMessagePic.Get().(*RefundMessagePic)
}

// ReleaseRefundMessagePic 释放RefundMessagePic
func ReleaseRefundMessagePic(v *RefundMessagePic) {
	v.Desc = ""
	v.Pic = ""
	v.PicToken = ""
	poolRefundMessagePic.Put(v)
}
